import berserk
from datetime import datetime, timezone
import json
import os
import sys
import time

TOK = os.environ.get("LICHESS_TOKEN")

sess = berserk.TokenSession(TOK)
client = berserk.Client(sess)

me = client.account.get()

# https://stackoverflow.com/a/22238613/42559
def json_serial(obj):
    """JSON serializer for objects not serializable by default json code"""

    if isinstance(obj, datetime):
        return millis_since_epoch(obj)
    raise TypeError("Type %s not serializable" % type(obj))


def millis_since_epoch(dt: datetime) -> int:
    return int(dt.timestamp() * 1000)


def datetime_from_seconds(ts):
    """Return the datetime for the given seconds since the epoch.

    UTC is assumed. The returned datetime is timezone aware.

    :return: timezone aware datetime
    :rtype: :class:`datetime`
    """
    return datetime.fromtimestamp(ts, timezone.utc)


def datetime_from_millis(millis):
    """Return the datetime for the given millis since the epoch.

    UTC is assumed. The returned datetime is timezone aware.

    :return: timezone aware datetime
    :rtype: :class:`datetime`
    """
    return datetime_from_seconds(millis / 1000)


if os.path.isfile("raw_games.json"):
    games = json.load(open("raw_games.json"))
else:
    outfile = open("raw_games.json", "w")
    games = list(
        client.games.export_by_player(
            "llimllib",
            max=60,
            moves=True,
            pgn_in_json=True,
            tags=True,
            clocks=True,
            evals=True,
            # export_by_player does not seem to support the `literate` flag. This is
            # probably a bug. https://github.com/ZackClements/berserk/pull/24
            # so I added it and installed my own version of the library
            literate=True,
            # sort the games from oldest to newest
            sort="dateAsc",
        )
    )
    json.dump(games, outfile, default=json_serial)

# The game stream is throttled, depending on who is making the request:
#
#     Anonymous request: 20 games per second
#     OAuth2 authenticated request: 30 games per second
#     Authenticated, downloading your own games: 60 games per second

# let's download 60 games per second and stuff them into a file - later we can
# process them intelligently into a database
while 1:
    try:
        since = millis_since_epoch(games[-1]["createdAt"])
    except AttributeError:
        # berserk returns times in datetime, but we save them as integers,
        # which is the input format that it's expecting. If createdAt isn't a
        # datetime, it might be a time we lodaed from the filesystem
        since = games[-1]["createdAt"]
    newgames = list(
        client.games.export_by_player(
            "llimllib",
            max=60,
            # we want to increase the "since" time by one since liches does >=,
            # and we don't want to re-download a game
            since=since + 1,
            moves=True,
            pgn_in_json=True,
            tags=True,
            clocks=True,
            evals=True,
            literate=True,
            # sort the games from oldest to newest
            sort="dateAsc",
        )
    )

    if not newgames:
        break

    games.extend(newgames)

    # rewrite our raw games file
    outfile = open("raw_games.json", "w")
    json.dump(games, outfile, default=json_serial)

    if len(set(g["id"] for g in games)) != len(games):
        print(
            "duplicates found, quitting", len(set(g["id"] for g in games)), len(games)
        )
        print(games[-1]["id"])
        sys.exit(64)

    time.sleep(1)
    sys.stdout.write(".")
    sys.stdout.flush()

print()

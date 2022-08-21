import json
import sqlite3
from typing import Any, List

DB = sqlite3.connect("chess.db")


def query(sql: str, *params: Any | None):
    c = DB.cursor()
    c.execute(sql, params)
    rows = c.fetchall()
    c.close()
    DB.commit()
    return rows


def dig(obj, *keys):
    """
    Return obj[key_1][key_2][...] for each key in keys, or None if any key
    in the chain is not found

    So, given this `obj`:

    {
        "banana": {
            "cream": "pie"
        }
    }

    dig(obj, "banana") -> {"cream": "pie"}
    dig(obj, "banana", "cream") -> "pie"
    dig(obj, "banana", "rama") -> None
    dig(obj, "Led", "Zeppelin") -> None
    """
    for key in keys:
        try:
            obj = obj[key]
        except (KeyError, IndexError):
            return None
    return obj


query(
    """
CREATE TABLE IF NOT EXISTS games(
    id text primary key,
    whitePlayerId text,
    whitePlayerName text,
    whitePlayerRating numeric,
    whitePlayerRatingDiff numeric,
    blackPlayerId text,
    blackPlayerName text,
    blackPlayerRating numeric,
    blackPlayerRatingDiff numeric,
    winner text,
    moves text,
    clockInitial numeric,
    clockIncrement numeric,
    clockTotal numeric,
    pgn text,
    rated boolean,
    variant text,
    speed text,
    perf text,
    createdAt numeric,
    lastMoveAt numeric,
    status text
)"""
)

games = json.load(open("raw_games.json"))
for game in games:
    try:
        query(
            """INSERT INTO games VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)""",
            game["id"],
            # some games don't have one player or the other - maybe if they
            # deleted their account or if they were caught cheating? Not sure
            dig(game, "players", "white", "user", "id"),
            dig(game, "players", "white", "user", "name"),
            dig(game, "players", "white", "rating"),
            dig(game, "players", "white", "ratingDiff"),
            dig(game, "players", "black", "user", "id"),
            dig(game, "players", "black", "user", "name"),
            dig(game, "players", "black", "rating"),
            dig(game, "players", "black", "ratingDiff"),
            # the winner field is not present in a draw. I'm a bit worried that
            # I'm calling some games a draw that aren't actually, but it'll do
            # for now
            game.get("winner", "drawn"),
            game["moves"],
            # some games don't have any clock; maybe games can be played with
            # the clock entirely off?
            dig(game, "clock", "initial"),
            dig(game, "clock", "increment"),
            dig(game, "clock", "totalTime"),
            game["pgn"],
            game["rated"],
            game["variant"],
            game["speed"],
            game["perf"],
            game["createdAt"],
            game["lastMoveAt"],
            game["status"],
        )
    except:
        print("failed to import", game)
        raise

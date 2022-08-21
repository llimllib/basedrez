package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
)

// getenv returns the environment variable given by the key if present,
// otherwise it returns adefault
func getenv(key string, adefault string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return adefault
	}
	return val
}

func initLogger() zerolog.Logger {
	debug := os.Getenv("DEBUG")

	if debug == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	var writer io.Writer
	if getenv("PRETTY_LOGGING", "") == "" {
		writer = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	} else {
		writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
	}
	logger := zerolog.New(writer).With().Timestamp().Str("component", "etl").Logger()
	return logger
}

func main() {
	log := initLogger()
	log.Info().Msg("starting server")

	host := getenv("HOST", "")
	port := getenv("PORT", "8000")
	addr := fmt.Sprintf("%s:%s", host, port)

	db, err := sql.Open("sqlite3", "../chess.db")
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to open database")
	}

	server := NewBasedrezServer(db, log)

	mux := http.NewServeMux()
	mux.HandleFunc("/", server.Index)

	srv := &http.Server{
		Addr:              addr,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           http.TimeoutHandler(mux, 2*time.Second, "request timed out"),
	}

	log.Info().Msgf("starting helloserver on %s", addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Error from server")
	}
}

type BasedrezServer struct {
	db        *sql.DB
	log       zerolog.Logger
	templates *template.Template
}

func NewBasedrezServer(db *sql.DB, log zerolog.Logger) *BasedrezServer {
	templates := template.Must(template.ParseGlob("./static/*.html"))
	server := BasedrezServer{
		db:        db,
		log:       log,
		templates: templates,
	}
	return &server
}

func (b BasedrezServer) mustExec(sqlc string) sql.Result {
	res, err := b.db.Exec(sqlc)
	if err != nil {
		b.log.Fatal().Err(err).Str("sql", sqlc).Msg("Unable to exec")
	}
	return res
}

func (b BasedrezServer) mustQuery(query string) *sql.Rows {
	res, err := b.db.Query(query)
	if err != nil {
		b.log.Fatal().Err(err).Str("sql", query).Msg("Unable to complete query")
	}
	return res
}

type Game struct {
	ID                string
	WhitePlayerID     string
	WhitePlayerName   string
	WhitePlayerRating string
	BlackPlayerID     string
	BlackPlayerName   string
	BlackPlayerRating string
	Winner            string
}

func (b BasedrezServer) RecentGames() []Game {
	rows := b.mustQuery(`
SELECT
	id,
	whitePlayerId, whitePlayerName, whitePlayerRating,
	blackPlayerId, blackPlayerName, blackPlayerRating,
	winner
FROM games
LIMIT 10`)

	var games []Game
	for rows.Next() {
		game := Game{}
		err := rows.Scan(
			&game.ID,
			&game.WhitePlayerID, &game.WhitePlayerName, &game.WhitePlayerRating,
			&game.BlackPlayerID, &game.BlackPlayerName, &game.BlackPlayerRating,
			&game.Winner)
		if err != nil {
			b.log.Fatal().Err(err).Msg("Unable to get recent games")
		}
		games = append(games, game)
	}

	return games
}

func (b BasedrezServer) Index(w http.ResponseWriter, r *http.Request) {
	idx := b.templates.Lookup("index.html")
	if idx == nil {
		b.log.Fatal().Msg("Failed looking up template")
	}

	games := b.RecentGames()
	b.log.Debug().Int("games", len(games)).Msg("# of games")
	if err := idx.Execute(w, struct{ Games []Game }{Games: games}); err != nil {
		b.log.Fatal().Err(err).Msg("Failed executing template")
	}
}

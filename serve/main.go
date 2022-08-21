package main

import (
	"database/sql"
	"fmt"
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

	server := BasedrezServer{
		db:  db,
		log: log,
	}

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
	db  *sql.DB
	log zerolog.Logger
}

func (b BasedrezServer) mustQuery(query string) sql.Result {
	res, err := b.db.Exec(query)
	if err != nil {
		b.log.Fatal().Err(err).Msg("Unable to complete query")
	}
	return res
}

func (b BasedrezServer) Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

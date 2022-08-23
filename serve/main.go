package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/llimllib/basedrez/ent"
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

	host := getenv("HOST", "localhost")
	port := getenv("PORT", "8000")
	addr := fmt.Sprintf("%s:%s", host, port)

	// client, err := ent.Open("sqlite3", "file:../chess.db?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open("sqlite3", "../chess.db")
	if err != nil {
		log.Fatal().Err(err).Msg("failed opening connection to sqlite")
	}
	// TODO: when do we want to run this? do we at all?
	// Run the auto migration tool.
	// if err := client.Schema.Create(context.Background()); err != nil {
	//     log.Fatalf("failed creating schema resources: %v", err)
	// }

	server := NewBasedrezServer(client, log)

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

	log.Info().Msgf("starting helloserver on http://%s", addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Error from server")
	}
}

type BasedrezServer struct {
	db        *ent.Client
	log       zerolog.Logger
	templates *template.Template
}

func NewBasedrezServer(db *ent.Client, log zerolog.Logger) *BasedrezServer {
	templates := template.Must(template.ParseGlob("./static/*.html"))
	server := BasedrezServer{
		db:        db,
		log:       log,
		templates: templates,
	}
	return &server
}

func (b BasedrezServer) Index(w http.ResponseWriter, r *http.Request) {
	idx := b.templates.Lookup("index.html")
	if idx == nil {
		b.log.Fatal().Msg("Failed looking up template")
	}

	games, err := b.db.Game.Query().Where().Limit(10).All(context.Background())
	if err != nil {
		b.log.Fatal().Err(err).Msg("Failed querying games")
	}

	if err := idx.Execute(w, struct{ Games []*ent.Game }{Games: games}); err != nil {
		b.log.Fatal().Err(err).Msg("Failed executing template")
	}
}

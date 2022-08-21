package main

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

func TestHelloServer(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/world", nil)
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	b := BasedrezServer{db: db, log: log.Logger}
	b.Index(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if !strings.Contains(string(data), "Welcome") {
		t.Fatalf("didn't find Welcome")
	}
}

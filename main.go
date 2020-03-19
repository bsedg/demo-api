package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var gitCommit string
var gitBranch string

const (
	API_PREFIX_V1 = "/api/v1"
)

type VersionDetails struct {
	Date       time.Time `json:"date"`
	APIVersion string    `json:"api_version"`
	GitCommit  string    `json:"git_commit"`
	GitBranch  string    `json:"git_branch"`
}

func newVersionDetails() *VersionDetails {
	return &VersionDetails{
		Date:       time.Now(),
		APIVersion: API_PREFIX_V1,
		GitCommit:  gitCommit,
		GitBranch:  gitBranch,
	}
}

func getVersionDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(newVersionDetails()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix(API_PREFIX_V1).Subrouter()
	api.Handle("/version", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(getVersionDetails)))

	log.Fatal(http.ListenAndServe(":8080", handlers.CompressHandler(api)))
}

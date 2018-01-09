package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const static = "dist"

func main() {
	r := newRouter(static)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func newRouter(staticPath string) *mux.Router {
	r := mux.NewRouter()
	staticDir := http.FileServer(http.Dir(staticPath))
	r.Handle("/", staticDir)
	r.PathPrefix("/{_:.*}").Handler(staticDir)
	return r
}

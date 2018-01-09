package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const static = "dist"

const (
	POST = http.MethodPost
)

func main() {
	r := newRouter(static)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func newRouter(staticPath string) *mux.Router {
	r := mux.NewRouter()
	staticDir := http.FileServer(http.Dir(staticPath))
	r.Handle("/", staticDir)
	r.PathPrefix("/{_:.*}").Handler(staticDir)
	r.HandleFunc("/fmt", goFmt).Methods(POST)
	r.HandleFunc("/runtests", runTests).Methods(POST)
	return r
}

func goFmt(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}
	defer r.Body.Close()
	log.Info(string(body))
}

func runTests(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}
	defer r.Body.Close()
	log.Info(string(body))
}

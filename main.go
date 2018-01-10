package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

const (
	static = "dist"
	port   = ":4000"
)

//Saves names of api methods to avoid using of magic strings
const (
	POST = http.MethodPost
	GET  = http.MethodGet
)

func main() {
	r := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "DELETE"},
	}).Handler(newRouter(static))

	log.Info("Server started on port 4000")
	log.Fatal(http.ListenAndServe(port, r))
}

func newRouter(staticPath string) *mux.Router {
	r := mux.NewRouter()
	staticDir := http.FileServer(http.Dir(staticPath))
	r.Handle("/", staticDir)
	r.Path("/fmt").HandlerFunc(goFmt).Methods(POST)
	r.Path("/runtests").HandlerFunc(runTests).Methods(POST)
	r.Handle("/{_:.*}", staticDir)
	return r
}

func goFmt(w http.ResponseWriter, r *http.Request) {
	log.Info("request for formatting is made")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}
	defer r.Body.Close()
	log.Info(string(body))
}

func runTests(w http.ResponseWriter, r *http.Request) {
	log.Info("request for running tests is made")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}
	defer r.Body.Close()
	log.Info(string(body))
	writeSolution(body)
}

func writeSolution(usercode []byte) {
	solutionFile, err := os.OpenFile("challenges/basic/solution.go", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer solutionFile.Close()

	n, err := solutionFile.Write(usercode)
	if err != nil || n != len(usercode) {
		log.Errorf("Failed to solution write to file: %s", err)
	}
	if err := solutionFile.Close(); err != nil {
		log.Fatal(err)
	}
}

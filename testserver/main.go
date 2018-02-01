package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"

	"github.com/pkg/errors"

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
	r.Path("/runtests/{testname}").HandlerFunc(runTests).Methods(POST)
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
	name := mux.Vars(r)["testname"]
	log.Info("request for running tests is made")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}
	defer r.Body.Close()
	log.Info(string(body))
	filepath := "challenges/basic/solution.go"
	if err = writeSolution(body, filepath); err != nil {
		log.Fatal(err)
	}
	res := execTest(name)

	log.Infof("%+v", res)
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func writeSolution(usercode []byte, filepath string) error {
	solutionFile, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer solutionFile.Close()

	n, err := solutionFile.Write(usercode)
	if err != nil || n != len(usercode) {
		return errors.Wrapf(err, "Failed to solution write to file: %s", filepath)
	}
	return nil
}

var re, _ = regexp.Compile("Expected {{(.*)}}, got {{(.*)}}")

func execTest(name string) *Result {
	prog := "go"
	args := []string{"test", "--run", name, "./challenges/basic/basic_test.go", "./challenges/basic/solution.go"}
	cmd := exec.Command(prog, args...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(err)
	}
	return parseOutput(stdoutStderr)
}

type Result struct {
	TestsPassed bool          `json:"passed"`
	BuildFailed bool          `json:"buildfailed"`
	BuildErr    []string      `json:"errors,omitempty"`
	FailedTests map[int]*Test `json:"failedtests,omitempty"`
}

type Test struct{ Expected, Got string }

func parseOutput(output []byte) *Result {
	result := &Result{}
	lines, buildfailed := checkBuild(output)
	log.Info(lines)
	if buildfailed {
		result.BuildFailed = true
		result.BuildErr = lines[1 : len(lines)-1]
		return result
	}

	if strings.Contains(lines[0], "FAIL") {
		result.FailedTests = make(map[int]*Test)
		log.Info("Tests failed")
		errors := lines[1 : len(lines)-2]
		log.Info("lines", errors)
		var wg sync.WaitGroup
		for i, e := range errors {
			wg.Add(1)
			go func(i int, t string) {
				result.FailedTests[i] = parseFailed(t)
				wg.Done()
			}(i, e)
		}
		wg.Wait()
	}

	if strings.Contains(lines[0], "PASS") {
		result.TestsPassed = true
	}

	return result
}

func checkBuild(out []byte) ([]string, bool) {
	output := string(out)
	lines := strings.Split(output, "\n")
	return lines[:len(lines)-1], strings.HasSuffix(output, "[build failed]")
}

func parseFailed(input string) *Test {
	match := re.FindStringSubmatch(input)
	if len(match) < 2 {
		return nil
	}
	return &Test{Expected: match[1], Got: match[2]}
}

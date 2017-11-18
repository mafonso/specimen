package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Uptime time.Duration `json:"uptime,omitempty"`
}

var status Status
var start time.Time

func rootHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	status.Uptime = t.Sub(start)
	json.NewEncoder(w).Encode(status)
}

func envHandler(w http.ResponseWriter, r *http.Request) {
	var env []string
	for _, pair := range os.Environ() {
		env = append(env, pair)
	}
	json.NewEncoder(w).Encode(env)
}

func handleRequests() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/env", envHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	start = time.Now()
	handleRequests()
}

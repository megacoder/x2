package main

import (
	"encoding/json"
	"fmt"
	"github.com/owainlewis/x2/agent"
	"log"
	"net/http"
)

var emily = agent.New()

func agentRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		query := agent.AgentQuery{}
		err := decoder.Decode(&query)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		reply := emily.Query(query)
		fmt.Fprintf(w, reply.Tell)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "static/index.html")
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {

	staticFileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static", staticFileServer)
	http.HandleFunc("/agent", agentRequestHandler)
	http.HandleFunc("/", indexHandler)

	log.Println("Starting agent...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

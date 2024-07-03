package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Initialize an in-memory store
var (
	//articles []Article
	mutex sync.Mutex
)

func main() {
	port := ":8080"
	fmt.Println("Hello word from Rocky API REST. Listen in" + port + " port")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			HelloWord(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			exampleRoute(w)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(port, nil)
}

func HelloWord(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello Word From Roky API REST.♥")
}

func exampleRoute(w http.ResponseWriter) {
	mutex.Lock()
	defer mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Ruta de ejemplo: example-route. Hola API Wategay in Kong :D")
}

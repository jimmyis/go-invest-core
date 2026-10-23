package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handleHealth)
	return mux
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func main() {
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", newMux()))
}

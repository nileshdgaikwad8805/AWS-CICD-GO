package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"status": "UP"})
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		writeJSON(w, map[string]string{
			"app":     "CloudPulse",
			"company": "VidyaOps",
			"version": "1.0.0",
			"message": "Deployed via AWS CodePipeline",
			"time":    time.Now().UTC().Format(time.RFC3339),
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // Elastic Beanstalk default
	}
	log.Println("CloudPulse listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

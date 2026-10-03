package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

const version = "2.0.0"

//go:embed static
var staticFiles embed.FS

var startedAt = time.Now()

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	json.NewEncoder(w).Encode(v)
}

func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"status": "UP"})
	})

	mux.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		host, _ := os.Hostname()
		env := os.Getenv("ENVIRONMENT")
		if env == "" {
			env = "production"
		}
		writeJSON(w, map[string]any{
			"app":           "CloudPulse",
			"company":       "VidyaOps",
			"version":       version,
			"goVersion":     runtime.Version(),
			"os":            runtime.GOOS + "/" + runtime.GOARCH,
			"hostname":      host,
			"environment":   env,
			"startedAt":     startedAt.UTC().Format(time.RFC3339),
			"uptimeSeconds": int64(time.Since(startedAt).Seconds()),
			"serverTime":    time.Now().UTC().Format(time.RFC3339),
			"message":       "Deployed via AWS CodePipeline",
		})
	})

	sub, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/", noCache(http.FileServer(http.FS(sub))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // Elastic Beanstalk default
	}
	log.Println("CloudPulse " + version + " listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

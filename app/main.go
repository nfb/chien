package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

var BINDADDR string = ":3000"

func catchall(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("meow"))
	slog.Info("Caught Request")
}

func interactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bodyb, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("Failed to read body: " + err.Error())
			http.Error(w, "Serverr", 500)
			return
		}
		body := string(bodyb)
		slog.Debug("Received Body: " + body)
		if body == "PING" {
			w.Write([]byte("PONG"))
			return
		}
	}
	fmt.Println(r)
	w.Write([]byte("meow"))
	slog.Info("Caught Request")
}

func configureLogging() {
	var programLevel = new(slog.LevelVar) // Info by default
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))

	switch loglevel := os.Getenv("LOGLEVEL"); loglevel {
	case "DEBUG":
		programLevel.Set(slog.LevelDebug)
	case "INFO":
		programLevel.Set(slog.LevelInfo)
	case "WARN":
		programLevel.Set(slog.LevelWarn)
	case "ERROR":
		programLevel.Set(slog.LevelError)
	default:
		programLevel.Set(slog.LevelInfo)
	}
}

func main() {
	envBindAddr = os.Getenv("BINDADDR")
	if envBindAddr != "" {
		BINDADDR = envBindAddr
	}
	fmt.Println("Starting server, binding to " + BINDADDR)
	http.HandleFunc("/*", catchall)
	http.HandleFunc("/api/interactions", interactionHandler)
	err := http.ListenAndServeTLS(BINDADDR, "cert.pem", "key.pem", nil)
	if err != nil {
		slog.Error("TLS Server returned error: " + err.Error())
	}

}

package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

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

func main() {
	fmt.Println("Starting...")
	http.HandleFunc("/*", catchall)
	http.HandleFunc("/api/interactions", interactionHandler)
	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	if err != nil {
		slog.Error("TLS Server returned error: " + err.Error())
	}
}

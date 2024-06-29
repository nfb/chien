package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func catchall(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("meow"))
	slog.Info("Caught Request")
}

func main() {
	fmt.Println("Starting...")
	http.HandleFunc("/*", catchall)
	err := http.ListenAndServeTLS(":3000", "cert.pem", "key.pem", nil)
	if err != nil {
		slog.Error("TLS Server returned error: " + err.Error())
	}
}

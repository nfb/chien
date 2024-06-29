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
	http.ListenAndServeTLS(":3000", "key,pem", "cert.pem", nil)
}

package main

import (
	"fmt"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "welcome to chun oj")
	if err != nil {
		return
	}
}

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: nil,
	}
	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}

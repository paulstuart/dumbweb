package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "0.0.0.0"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "[%-6s] %q\n", r.Method, html.EscapeString(r.URL.Path))
	})

	listen := addr + ":" + port
	log.Println("listening on: ", listen)

	log.Fatal(http.ListenAndServe(listen, nil))
}

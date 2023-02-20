package main

import (
	"embed"
	"flag"
	"fmt"
	"html"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	//go:embed static/*
	content embed.FS
)

func main() {
	var maintenance bool
	flag.BoolVar(&maintenance, "maintenance", false, "show maintenance page")
	flag.Parse()

	if ok, _ := strconv.ParseBool(os.Getenv("MAINTENANCE")); ok {
		maintenance = true
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "0.0.0.0"
	}

	if maintenance {
		log.Println("MAINTAIN!")
		dir, err := fs.Sub(content, "static")
		if err != nil {
			log.Fatalf("failed to get static dir: %v", err)
		}
		http.Handle("/", http.FileServer(http.FS(dir)))
		// http.Handle("/", http.FileServer(http.FS(content)))
		// http.Handle("/", http.StripPrefix("/static/", http.FileServer(http.FS(content))))
		// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(content))))
	} else {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "[%-6s] %q\n", r.Method, html.EscapeString(r.URL.Path))
		})
	}

	listen := addr + ":" + port
	log.Println("listening on: ", listen)

	log.Fatal(http.ListenAndServe(listen, nil))
}

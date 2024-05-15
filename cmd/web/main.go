package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// command line flags
	addr := flag.String("addr", ":3000", "HTTP network address")
	flag.Parse()

	// instantiate servers
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// register routes
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// start the server
	log.Printf("starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}

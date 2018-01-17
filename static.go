package main

import (
	"log"
	"net/http"

	"github.com/nogoegst/static/noindexfs"
)

func main() {
	fs := noindexfs.New(http.Dir("."))
	err := http.ListenAndServe(":80", http.FileServer(fs))
	if err != nil {
		log.Fatal(err)
	}
}

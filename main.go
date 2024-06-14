package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("views", "index.html"))
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views"))))

	fmt.Println("starting server on port :8080 ")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("error: %v", err)
	}
}

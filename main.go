package main

import (
	"context"
	"fmt"
	"fuuuuuuuuuk/views"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/send", postSend)

	fmt.Println("starting server on port :8080 ")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("error: %v", err)
	}
}

func postSend(w http.ResponseWriter, r *http.Request) {
	val := r.PostFormValue("send")
	log.Println(val)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := views.Body().Render(context.Background(), w)
	if err != nil {
		fmt.Println(err)
	}

}

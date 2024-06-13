package main

import (
	"fmt"
	"fuuuuuuuuuk/cmd/web/hanlders"
	"log"
	"net/http"
)

const PORT string = ":4030"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", hanlders.Landing)

	fmt.Println("server runnin")
	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		log.Fatal("cannot get server running: ", err)
		return
	}
}

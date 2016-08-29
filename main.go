package main

import (
	"github.com/nirasan/go-kaleidoscope/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/mono", handler.Mono)
	http.HandleFunc("/stripe", handler.Stripe)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"github.com/nirasan/go-kaleidoscope/handler"
	"github.com/nirasan/go-kaleidoscope/handler/kaleidoscope"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", kaleidoscope.Kaleidoscope)
	http.HandleFunc("/mono", handler.Mono)
	http.HandleFunc("/stripe", handler.Stripe)
	http.HandleFunc("/rotate", handler.Rotate)
	http.HandleFunc("/sector", handler.Sector)
	http.HandleFunc("/sector2", handler.Sector2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"github.com/StephenGriese/hello-api/handlers"
	"log"
	"net/http"

	"github.com/StephenGriese/hello-api/handlers/rest"
)

func main() {

	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/roster", rest.RosterHandler)

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

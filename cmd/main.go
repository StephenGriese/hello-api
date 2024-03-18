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

	mux.HandleFunc("/translate/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

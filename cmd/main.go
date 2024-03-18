package main

import (
	"fmt"
	"github.com/StephenGriese/hello-api/handlers"
	"log"
	"net/http"
	"os"

	"github.com/StephenGriese/hello-api/handlers/rest"
)

func main() {

	addr := fmt.Sprintf(":$s", os.Getenv("PORT"), "error")
	if addr == ":" {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/roster", rest.RosterHandler)

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

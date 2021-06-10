package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", ivyHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("No PORT provided, defaulting to port %v", port)
	}

	log.Printf("Listening on port %v", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

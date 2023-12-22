package main

import (
	"fmt"
	"net/http"
)

func main() {
	healthCheckHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Service is healthy!")
	}

	http.HandleFunc("/health", healthCheckHandler)

	port := 8080
	fmt.Printf("Application is running on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

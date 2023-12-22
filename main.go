package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	version string
	build   string
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	port := getEnv("APP_PORT", "8080")

	healthCheckHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Service is healthy!")
	}

	http.HandleFunc("/health", healthCheckHandler)

	fmt.Printf("Application is running on port:%s Version:%s Build:%s\n", port, version, build)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

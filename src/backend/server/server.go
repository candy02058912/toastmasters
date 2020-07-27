package main

import (
	"fmt"
	"log"
	"os"
	"svr"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var (
	port      = getEnv("PORT", "8080")
	toastType = getEnv("TYPE", "plain")
)

func main() {
	fmt.Println("start server")

	server := svr.NewServer(port, toastType)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

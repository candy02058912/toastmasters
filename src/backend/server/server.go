package main

import (
	"fmt"
	"log"
	"os"
	"svr"
	"time"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var (
	port     = getEnv("PORT", "8080")
	waitTime = getEnv("WAIT_TIME", "1000ms")
)

func main() {
	fmt.Println("start server")
	wt, err := time.ParseDuration(waitTime)
	if err != nil {
		log.Fatalf("error time duration format %s: %v", waitTime, err)
	}

	server := svr.NewServer(port, wt)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

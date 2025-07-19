package main

import (
    "fmt"
    "os"
)

var baseDomain = "example.com" // env BASE_DOMAIN

func main() {
    if envDomain := os.Getenv("BASE_DOMAIN"); envDomain != "" {
        baseDomain = envDomain
    }

    if err := startHTTPServer(":8080"); err != nil {
        fmt.Println("Error starting HTTP server:", err)
        os.Exit(1)
    }
}
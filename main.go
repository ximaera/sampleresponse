package main

import (
    "fmt"
    "os"
)

var baseDomain = "example.com" // env BASE_DOMAIN
var httpPort = "8080"          // env HTTP_PORT

func main() {
    if envDomain := os.Getenv("BASE_DOMAIN"); envDomain != "" {
        baseDomain = envDomain
    }

    if envHttpPort := os.Getenv("HTTP_PORT"); envHttpPort != "" {
        httpPort = envHttpPort
    }

    if err := startHTTPServer(":" + httpPort); err != nil {
        fmt.Println("Error starting HTTP server:", err)
        os.Exit(1)
    }
}
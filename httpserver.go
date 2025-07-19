package main

import (
    "errors"
    "fmt"
    "net"
    "net/http"
    "strconv"
    "strings"
    "time"
)

func startHTTPServer(addr string) error {
    http.HandleFunc("/", httpHandler)

    fmt.Printf("HTTP server listening on %s...\n", addr)
    return http.ListenAndServe(addr, nil)
}

func returnError(w http.ResponseWriter, msg string, status int) error {
    http.Error(w, msg, status)
    return errors.New(msg)
}

func sanityCheck(w http.ResponseWriter, r *http.Request) (int, string, error) {
    host := r.Host

    // Remove port from host if present
    if strings.Contains(host, ":") {
        var err error
        host, _, err = net.SplitHostPort(host)
        if err != nil {
            return 0, "", returnError(w, "Invalid Host header", http.StatusBadRequest)
        }
    }

    // Must end with base domain
    if !strings.HasSuffix(host, baseDomain) {
        return 0, "", returnError(w, "Forbidden", http.StatusForbidden)
    }

    // Parse subdomain: N.status.http.example.com → ["N", "status", "http"]
    parts := strings.Split(strings.TrimSuffix(host, "."+baseDomain), ".")
    if len(parts) < 3 {
        return 0, "", returnError(w, "Bad request", http.StatusBadRequest)
    }

    codeStr, mode := parts[0], parts[1]
    value, err := strconv.Atoi(codeStr)
    if err != nil || value < 0 {
        return 0, "", returnError(w, "Invalid numeric value", http.StatusBadRequest)
    }

    return value, mode, nil
}

func processStatus(w http.ResponseWriter, code int) {
    if code < 100 || code > 999 {
        http.Error(w, "Invalid HTTP status code", http.StatusBadRequest)
        return
    }
    w.WriteHeader(code)
    fmt.Fprintf(w, "Returned HTTP status code %d\n", code)
}

func processTimeout(w http.ResponseWriter, timeout int) {
    hj, ok := w.(http.Hijacker)
    if !ok {
        http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
        return
    }
    conn, _, err := hj.Hijack()
    if err != nil {
        return
    }
    go func() {
        defer conn.Close()
        time.Sleep(time.Duration(timeout) * time.Second)
    }()
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
    var value, mode, err = sanityCheck(w, r)
    if err != nil {
        return
    }
    
    switch mode {
    case "status":
        processStatus(w, value)
    
    case "timeout":
        processTimeout(w, value)
    
    default:
        http.Error(w, "Unknown mode", http.StatusBadRequest)
    }
}
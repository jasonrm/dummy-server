package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
    http.HandleFunc("/", handleRequest)
	fmt.Println(time.Now().Format(time.RFC3339), "Ready");
    http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	requestedStatusString := r.URL.Path[1:]
    if requestedStatusString == "" {
        requestedStatusString = "200"
    }
    status, _ := strconv.Atoi(requestedStatusString)
	w.WriteHeader(status)
	fmt.Fprint(w, status)
	fmt.Println(time.Now().Format(time.RFC3339), r.Method, r.URL, r.UserAgent())
}

package main

import (
    "fmt"
    "net/http"
    "strconv"
    "time"
    "flag"
)

func main() {
    portPtr := flag.Int("port", 8080, "port number")
    flag.Parse()

    listenStr := fmt.Sprintf(":%d", *portPtr)

    fmt.Println(time.Now().Format(time.RFC3339), "Listening On", listenStr);
    http.HandleFunc("/", handleRequest)
    http.ListenAndServe(listenStr, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	requestedStatusString := r.URL.Path[1:]
    if requestedStatusString == "" {
        requestedStatusString = "200"
    }
    status, err := strconv.Atoi(requestedStatusString)
    if err != nil {
        status = 200
    }
	w.WriteHeader(status)
	fmt.Fprint(w, status)
	fmt.Println(time.Now().Format(time.RFC3339), r.Method, status, r.URL, r.UserAgent())
}

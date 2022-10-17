package main

import (
	"fmt"
	"net/http"
	"time"
)

func timestamp(now time.Time) string {
	rv := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	return (rv)
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	qry := r.URL.Query()
	thing := qry.Get("thing")
	fmt.Fprintf(w, "[%s] - Go Web - Hello\nPath  = %s\nQuery = %s\nthing = %s\n", timestamp(now), r.URL, qry, thing)
}

func main() {
	http.HandleFunc("/", reqHandler)
	http.ListenAndServe(":9999", nil)
}

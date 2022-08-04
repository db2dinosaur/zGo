package main

import (
  "fmt"
  "net/http"
  "time"
)

func timestamp(now time.Time) string {
  rv := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
  return(rv)
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
  now := time.Now()
  qry := r.URL.Query()
  pan := qry.Get("pan")
  fmt.Fprintf(w, "[%s] - Go Web - Hello\nPath  = %s\nQuery = %s\npan   = %s\n",timestamp(now),r.URL,qry,pan)
}

func main() {
  http.HandleFunc("/",reqHandler)
  http.ListenAndServeTLS(":9999", "cert.pem", "key.pem", nil)
}

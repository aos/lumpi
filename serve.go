package main

import (
  "fmt"
  "log"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there -- This is my go server!")
}

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":1234", nil))
}

package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)

type Message struct {
  Message string `json:"message"`
}

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.URL.Path);
  fmt.Fprintf(w, "Welcome, %1", r.URL.Path[1:])
}

func about(w http.ResponseWriter, r *http.Request) {
  m := Message{"Welcome to this terrible API, build v0.0.1"}
  b, err := json.Marshal(m)

  if err != nil {
    panic(err)
  }

  w.Write(b);
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/about", about)
  http.ListenAndServe(":8080", nil)
}

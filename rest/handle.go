package main

import (
  "fmt"
  "log"
  "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Println("hello function")
  fmt.Fprintf(w, "hello")
}

func test(w http.ResponseWriter, r *http.Request) {
  fmt.Println("/test")
  fmt.Fprintf(w, "test")
}


func main() {
  http.HandleFunc("/", hello)
  http.HandleFunc("/test", test)
  if err := http.ListenAndServe(":9090", nil); err != nil {
    log.Fatal("ListenAndServe", err)
  }
}


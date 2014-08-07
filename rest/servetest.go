package main

import(
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi thre, i love %s", r.URL.Path[1:])
}

func main() {
  // if user accesse for "/", then server redirect to handler function
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}


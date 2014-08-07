package main

import "net/http"
import "fmt"
func main() {
  resp, err := http.Get("http://google.com/")
  
  if err != nil {
    panic(err)
  }

  fmt.Println("response is below")
  fmt.Println(resp)


}



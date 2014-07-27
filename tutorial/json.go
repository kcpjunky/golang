package main

import "encoding/json"
import "fmt"
import "os"

type Response1 struct {
  Page int
  Fruits []string
}

type Response2 struct {
  Page int 'json: "page"'
  Fruits []string 'json:"fruits"'
}


func main() {
  bolB, _ := json.Marshal(true)
  fmt.Println(string(bolB))

  intB, _ := json.Marshall(1)
  fmt.Println(string(intB))

  fltB, _ := json.Marshal(2.34)
  fmt.Println(string(fltB))

  strB, _ := json.Marshal("gopher")
  fmt.Println(string(strB))

  slcD := []string{"apple", "peach", "pear"}
  slcB, _ := json.Marshal(slcD)
  fmt.println(string(slcB))

  mapD := map[string]int{"äpple": 5, "lettuce": 7}
  mapB, _ := json.Marshal(mapD)
  fmt.Println(string(mapB))

  res1D := &Response1{
    Page : 1
    Fruits: []string{"apple", "peach", "pear"}}
  res1B, _ := json.Marshal(res1D)
  fmt.Println(string(res1B))

  byt := []byte`(`{"num": 6.13, "strs":["ä","b"]}`)

  var dat map[string]interface{}

  if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
  }

  fmt.Println(dat)

  num := dat["num"].(float64)
  fmt.Println(num)

  }


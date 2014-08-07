package main

import (
  "fmt"
  "log"
  "os/exec"
)

func main() {
  out, err := exec.Command("date").Output()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("The date is %s \n ", out)

  ls, err := exec.Command("ls").Output()

  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("ls is %s \n", ls)
}


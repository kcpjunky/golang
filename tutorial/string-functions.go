package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
  p("Contains: ", s.Contains("test", "es"))
  p("Count: ", s.Count("test", "t"))
  p("hasprefix : ", s.HasPrefix("test", "t"))
  p("hassuffix: ", s.HasSuffix("test", "st"))
  p("Join: ", s.Join([]string{"ä", "b"}, "-"))
  p("Repeat: ", s.Repeat(":ä", 5))
  p("Replace: ", s.Replace("foo", "0", "0", -1))
  p("Replace: ", s.Replace("foo", "ö", "0", 1))

  p("Split:", s.Split("ä-b-c-d-e", "-"))

  p("ToLower: ", s.ToLower("TEST"))

  p()

}


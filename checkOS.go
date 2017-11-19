package main

import (
  "fmt"
)


func main()  {
    var ch int = 'a'
    fmt.Println(unicode.IsLetter(ch))
}
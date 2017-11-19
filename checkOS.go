package main

import (
  "fmt"
  "strings"
)


func main()  {
  // str := "Beginning of the string " + 
  //     "continuation of the string"
  // s := "hel" + "lo"
  // s += " world"
  // var str string = "This is just an example of a string"
  // fmt.Printf("T/F? does the string \"%s\" has prefix %s?", str, "Th")
  // fmt.Printf("%t\n", strings.hasPrefix(str, "Th"))
  var str string = "Hi, I am Adam, Hi"
  fmt.Printf("The position of \"marc\" is: ")
  fmt.Printf("%d\n", strings.Index(str, "Marc"))
}
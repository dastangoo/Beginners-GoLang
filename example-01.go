package main
import "fmt"

func main()  {
  var value int
  var isPresent bool
  
  
  map1 := make(map[string]int)
  map1["New Delhi"] = 55
  map1["Mumbai"] = 20
  map1["China"] = 25
  value, isPresent = map1["Mumbai"]
  
  if isPresent {
    fmt.Printf("The value of \"Mumbai\" in map1: %d\n", value)
  } else {
    fmt.Println("map1 does not contain Mumbai")
  }
  
  value, isPresent = map1["Paris"]
  fmt.Printf("Is \"Paris\" in map1 ?: %t\n", isPresent)
  fmt.Printf("Value is: %d\n", value)
  
  delete(map1, "China")
  value, isPresent = map1["China"]
  if isPresent {
    fmt.Printf("The value of \"China\" in map1 is: %d\n", value)
  } else {
    fmt.Println("map1 does not contain \"China\"")
  }
}
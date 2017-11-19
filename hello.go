package main
import (
  "fmt"
)
func main()  {
  // path := os.Getenv("PATH")
  // fmt.Printf("PATH is: %s\n", path);  
  // a := 50
  // b := false
  // a = 20
  // var a string = "abc"
  // var a b c = 2, 9 "abc"
  a := 5
  b := 10
  a, b = b, a
  
  fmt.Println(a, b)
}


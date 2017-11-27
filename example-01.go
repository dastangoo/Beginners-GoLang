package main
import (
  "fmt"
  "./hello"
)

func main()  {
  var test1 string
  test1 = pack1.ReturnStr()
  fmt.Printf("ReturnStr from package1: %s\n", test1)
  fmt.Printf("Integer from package1: %d\n", pack1.pack1Int)
}
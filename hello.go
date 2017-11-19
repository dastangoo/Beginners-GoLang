package main
import (
  "fmt"
  "os"
)
func main()  {
  // var a, b *int
  // var a int
  // var b bool
  // var str string

  // var (
  //   a int
  //   b bool
  //   str string
  // )
  
  // var a int = 15
  // var i = 5
  // var b bool = false
  // var str string = "Go says hello to the world"
  
  // var a = 15
  // var b = false
  // var str = "Go says hello to the world"

  // var (
  //   a = 15
  //   b = false
  //   str = "Go says hello to the world"
  //   numShips = 50
  //   city string
  // )
  
  // var (
  //   HOME = os.Getenv("HOME")
  //   USER = os.Getenv("USER")
  //   GOROOT = os.Getenv("GOROOT")
  // )
  
  // var goos string = os.Getenv("GOOS")
  path := os.Getenv("PATH")
  fmt.Printf("PATH is: %s\n", path);  
}


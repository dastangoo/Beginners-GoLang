package main 
import "fmt"

func main() {
  var first int = 10
  var cond int 
  if first <= 0 {
    fmt.Println("first is less than or equal to 0\n")
  } else if first > 0 && first <5 {
    fmt.Println("first is between 0 and 5\n")
  } else {
    fmt.Println("first is greater or equal to 5\n")
  }
  
  if cond = 5; cond > 10 {
    fmt.Println("cond is is greater than 10\n")
  } else {
    fmt.Println("cond is not greater than 10\n")
  }
}
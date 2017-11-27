package main
import "fmt"

func main()  {
  var arrAge = [5]int{18, 20, 15, 22, 16}
  var arrLazy = [...]int{5, 6, 7, 8, 22}
  fmt.Println(len(arrAge))
  fmt.Println(len(arrLazy))
  var arrKeyValue = [5]string{0: "Adam", 1: "Eve" }
  
  for i := 0; i < len(arrKeyValue); i++ {
    fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
  }
}
package main
import "fmt"

func main()  {
  Function1()
}

func Function1()  {
  fmt.Printf("In Function1 at top\n")
  defer Function2()
  fmt.Printf("In Function1 at bottom\n")
}

func Function2()  {
  fmt.Printf("Function2: Deferred until the end of calling function!")
}
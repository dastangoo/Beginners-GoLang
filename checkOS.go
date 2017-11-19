package main
import (
  "fmt"
  "./trans"
)

var twoPi = 2 * trans.Pi

func main()  {
  fmt.Printf("2*Pi = %g/n", twoPi) // 2*pi = 6.28
}

func init()  {
  // setup the preparation
  go backend()
}
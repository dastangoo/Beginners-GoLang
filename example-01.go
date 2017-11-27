package main
import "fmt"

func main()  {
  s := "\u00ff\u754c"
  fmt.Println(s)
  for i, c := range s {
    fmt.Printf("%d:%c", i, c)
  } 
}
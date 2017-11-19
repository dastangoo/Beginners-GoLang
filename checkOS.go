package main

import (
  "fmt"
  // "time"
  // "rand"
)


func main()  {
  // a , b := 1, 2
  // 
  // b += a
  // b -= a
  // b *= a
  // b /= a
  // b %= a
  // b++
  // b--
  // fmt.Println(9/4)
  // fmt.Println(9%4)
  // fmt.Println(a)
  // fmt.Println(b)
  // for i := 0; i < 10; i++ {
  //   a := rand.Int()
  //   fmt.Printf("%d /", a)
  // }
  // 
  // for i := 0; i < 5; i++ {
  //   r := rand.Intn(8)
  //   fmt.Printf("%d /", r)
  // }
  // 
  // fmt.Printf()
  // timens := int64(time.Now().Nano)
  // rand.Seed(timens)
  // for i := 0; i  < 10; i++ {
  //   fmt.Printf("%2.2f /", 100*rand.Float32())
  // }
  
  type TZ int
  var a, b TZ = 3, 4
  c := a + b
  fmt.Printf("c has the value: %d", c) // prints c has value 7
}
package main

import (
  "fmt"
  "math"
)


func Uint8FromInt(n int) (uint8, error) {
  if 0 <= n && n <= math.MaxUint8 { // conversion is safe
    return uint8(n), nil
  }
  return 0, fmt.Errorf("%d is out of the uint8 range", n)
}
func IntFromFloat64(x float64) int {
  if math.MinInt32 <= x && x <= math.MaxInt32 { // x lies in the integer range
    whole, fraction := math.Modf(x)
    if fraction >= 0.5 {
      whole++
    }
    return int(whole)
  }
  panic(fmt.Sprintf("%g is out of int32 range", x))
}
func main()  {
    // fmt.Println(Uint8FromInt(100000000000))
    // fmt.Println(IntFromFloat64(65.9))
    var c1 complex64 = 5 + 10i
    fmt.Printf("The value is: %v", c1)
}
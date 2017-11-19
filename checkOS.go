package main

import (
  "fmt"
)


func main()  {
  type ByteSize float64
  const (
    _ = iota // ignore the first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
  )
  type BitFlag int
  const (
    Active BitFlag = 1 << iota // 1 << 0 == 1
    Send                       // 1 << 1 == 2
    Receive                    // 1 << 2 == 4
  )
  flag := Active | Send
  fmt.Println(flag)
}
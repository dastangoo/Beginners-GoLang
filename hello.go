package main
import "fmt"
func main()  {
  // const beef, two, c = "meat", 2, "veg"
  // const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
  // 
  // 
  // const (
  //   Monday, Tuesday, Wednesday = 1, 2, 3
  //   Thursday, Friday, Saturday = 4, 5, 6
  // )
  // 
  // const (
  //   unknown = 0
  //   Female = 1
  //   Male = 2
  // )
  // 
  // const (
  //   a = iota
  //   b 
  //   c
  // )
  // 
  // const (
  //   Sunday = iota
  //   Monday
  //   Tuesday
  //   Wednesday
  //   Thursday
  //   Friday
  //   Saturday
  // )

  type Color int 
  const (
    RED Color = iota // 0
    ORANGE // 1
    YELLOW // 2
    GREEN // 3
    BLUE // 4
    INDIGO // 5
    VIOLET // 6
  )
  fmt.Println(YELLOW);  
}


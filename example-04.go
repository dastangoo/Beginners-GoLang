package main
import "fmt"

func main()  {
  doDBOperations()
}

func connectToDB()  {
  fmt.Println("ok, connected to db")
}

func disconnectFromDB()  {
  fmt.Println("ok, disconnected from db")
}

func doDBOperations()  {
  connectToDB()
  fmt.Println("Defering the database disconnect.")
  defer disconnectFromDB() // function called with defer
  fmt.Println("Doing some DB operations.")
  fmt.Println("Oops!")
  fmt.Println("Oops! some crash or network error has occured")
  fmt.Println("Returing from function here")
  return // terminate the program
  // deferred function executed here
}
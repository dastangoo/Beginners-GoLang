package main

import (
  "fmt"
  "strings" 
)


func main() {
    // var str string = "hi, I am Adam, Hi"
    // fmt.Printf("The position of \"marc\" is: ")
    // fmt.Printf("%d\n", strings.Index(str, "Marc"))
    // fmt.Printf("The position of the first instance of \"hi\" is: ")
    // fmt.Printf("%d\n", strings.Index(str, "Hi"))
    // fmt.Printf("The position of the last instance of \"hi\" is: ")
    // fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))
    // fmt.Printf("The position of \"Burger\" is: ")
    // fmt.Printf("%d\n", strings.Index(str, "Burger"))    
    
    
    // strings.Replace(str, old, new, n)
    strings.Replace("Hi, I am Adam, hi", "Hi", "Bye", 2)
    
    var str string = "hello, how is it going, Adam?"
    var manyG = "ggggggg"
    
    fmt.Printf("Number of H's in %s: ", str)
    fmt.Printf("%d\n", strings.Count(str, "H"))
    
    fmt.Printf("Number of double g's in %s is: ", manyG)
    fmt.Printf("%d\n", strings.Count(manyG, "gg"))
    
}
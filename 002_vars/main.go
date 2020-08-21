package main

import "fmt"

var name string = "Bruce"
func main(){
  // Using the var keyword
  //var name string = "Bruce"
  lastName := "Gawade" // shorthand assignmment
  var age int = 37
  const isCool = true  // You cannot assign new values to constant
  size := 1.3
  // name, email := "aditya", "aditya2706@gmail.com" --> You can do multiple assignments as well
  fmt.Println(name, age, isCool, lastName, size)
  fmt.Printf("%T\n", name) // to get type of the variable
  fmt.Printf("%T\n", age)
  fmt.Printf("%T\n", isCool)
  fmt.Printf("%T\n", lastName)
  fmt.Printf("%T\n", size)
  
}

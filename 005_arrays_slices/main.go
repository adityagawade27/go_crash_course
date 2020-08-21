package main

import "fmt"

func main(){
  /*
  // Arrays
  var fruitArr1 [2]string

  // Assign values
  fruitArr1[0] = "Apple"
  fruitArr1[1] = "Orange"

  // Declare and assign
  fruitArr2 := [2]string{"Mango", "Banana"}
  fmt.Println(fruitArr)
  fmt.Println(fruitArr[1])
  */

  // Slices

  fruitSlice := []string{"Mango", "Banana", "Grape", "Cherry"}
  fmt.Println(fruitSlice)

  // to count the no of elements
  fmt.Println(len(fruitSlice))
  fmt.Println(fruitSlice[1:2])
  fmt.Println(fruitSlice[1:3])


}

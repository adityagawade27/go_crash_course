// Interfaces can be used to define generalized functions that take various inputs
// Eg: interface to caluclate area ; one interface can be use to calculate are of
// rectangle and circle instead of using multiple functions

package main

import (
  "fmt"
  "math"
)

// Defione interface
type Shape interface {
  area() float64
}

type Circle struct {
  x, y, radius float64
}

type Rectangle struct {
  width, height float64
}

// Value reciever
func (c Circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

// Value reciever
func (r Rectangle) area() float64 {
    return r.width * r.height
}

func getArea(s Shape) float64 {
  return s.area()
}

// This avoids having to classify shape and then call each function indivually
// Instead you can juyst do a call to the interface basede function and let it
// calculate the area based on the struct sent


func main(){
  fmt.Println("Hello World")
  circle := Circle{x:0, y:0, radius:5}
  rectangle  := Rectangle{width:10, height:5}

  fmt.Printf("Circle Area: %f\n", getArea(circle))
  fmt.Printf("Circle Area: %f\n", getArea(rectangle))
}

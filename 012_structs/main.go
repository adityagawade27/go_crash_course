package main

import (
  "fmt"
  "strconv"

)


// Define Person struct

type Person struct {
  firstName string
  lastName string
  city string
  gender string
  age int
}

// Value reciever method -- Greeting
// Use value reciever when you want to do calculations and not actually change the data
 func (p Person) greet() string {
   return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
 }



// Pointer reciever method  -- hasBirthday
// Use pointer reciever when you want to actually change values in the struct
func(p *Person) hasBirthday() {
  p.age++
}

// getMarried  --> pointer reciever
func(p *Person) getMarried(spouseLastName string) {
   if p.gender == "m" {
     return
   } else {
     p.lastName = spouseLastName
   }
  }

func main(){
  // Init person using struct
  person1 := Person{firstName: "Bruce", lastName: "Wayne", city:"Gotham", gender:"m", age:38}
  // Without using property names
  person2 := Person{"Lois","Lane","Metropolis","f",32}

  fmt.Println(person1)
  fmt.Println(person2)
  fmt.Println(person1.firstName)
  person1.age++
  fmt.Println(person1)

  person1.hasBirthday()
  fmt.Println(person1.greet())

 person2.getMarried("Wayne")
 fmt.Println(person2.greet())

 person1.getMarried("Kent")
 fmt.Println(person1.greet())
}

package main

import "fmt"

func main(){
  // Define Map
  emails := make(map[string]string)

  // Assign kv
  emails["Bruce"] = "bruce@batman.com"
  emails["Selena"] = "selena@catwoman.com"
  emails["Clark"] = "clark@superman.com"

  fmt.Println(emails)
  fmt.Println(emails["Bruce"])
  fmt.Println(len(emails))

  // Delete from Map
  delete(emails, "Bruce")
  fmt.Println(emails)

  /*
  Declare map and aded kv
  emails := map[string]string{"Bruce": "bruce@batman.com", "Selena": "selena@catwoman.com"} 
  */
}

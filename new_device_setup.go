package main

import (
  "fmt"
)

func main() {
  var name string
  fmt.Println("What is your name?")
  fmt.Scan(&name)
  fmt.Println("Hello", name)
  var age int
  fmt.Println("What is your age?")
  fmt.Scan(&age)
  fmt.Printf("%s is %d years old.\n", name, age)
  newMessage := fmt.Sprintf("Hi I'm %s and I am %d years old.", name, age)
  fmt.Println(newMessage)
}

package main

import "fmt"

func main() {
  fmt.Println("Hello, world.")

  var whatToSay string
  var i int

  whatToSay = "This is the day I learn Golang!"

  fmt.Println(whatToSay)

  i = 7

  fmt.Println("i is set to", i)

  saySomething()
}

func saySomething() string {
  return "something"
}

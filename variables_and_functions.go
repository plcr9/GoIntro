package main

import "fmt"

var myName string

func main() {
  fmt.Println("Hello, world.")

  var whatToSay string
  var i int

  whatToSay = "This is the day I learn Golang!"

  fmt.Println(whatToSay)

  i = 7

  fmt.Println("i is set to", i)

  whatWasSaid := saySomething()

  fmt.Println("The function returned", whatWasSaid)
}

func saySomething() string {
  return "something"
}

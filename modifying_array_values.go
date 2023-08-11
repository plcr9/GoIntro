package main

import "fmt"

func main() {
  myDogs := [3]string{"Frida", "Fedo", "Jegf"}
  fmt.Println(myDogs)
  myDogs[1] = "Fido"
  myDogs[2] = "Jeff"
  fmt.Println(myDogs)
}

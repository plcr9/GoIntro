package main

import (
  "fmt"
)

func askOrder() (string) {
  var input string
  fmt.Println("What would you like to eat: ")
  fmt.Scan(&input)
  return input
}

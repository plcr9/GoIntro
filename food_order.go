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

func contains(menu []string, order string) bool {
  for _, item := range menu {
    if order == item {
      return true
    }
  }
  return false
}

func main() {
  fastfoodMenu := []string("Burgers", "Nuggets", "Fries"}
  fmt.Println("The fast food menu has these items:", fastfoodMenu)

  var total int
  var order string

  for order != "quit" {
    order = askOrder()
  }

  fmt.Println("The total for the order is", total)
}


                           

                      

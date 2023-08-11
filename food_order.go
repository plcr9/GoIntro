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
    if contains(fastfoodMenu, order) {
      total+=4
  } else {
      fmt.Println("This item is not on the menu.")
  }
}

  for amount := total; amount > 0; amount -= 2 {
    fmt.Println("The amount owed is:", amount)
    fmt.Println("Handing over a two dollar bill.")
  }
                           
  fmt.Println("The total for the order is", total)
}


                           

                      

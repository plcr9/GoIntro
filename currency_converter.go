package main

import (
  "fmt"
)

func main() {
  currencies := map[string]float32{"EUR":0.95, "JPY":130}

  var dollarAmount float32
  var currency string
  fmt.Println("Please enter a dollar amount:")
  fmt.Scan(&dollarAmount)
  if dollarAmount == 0 {
    fmt.Println("Invalid dollar amount.")
  } else {
    fmt.Println("Please enter a currency (EUR or JPY):")
}

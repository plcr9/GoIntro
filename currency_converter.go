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
}

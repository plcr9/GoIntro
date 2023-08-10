package main

import "fmt"

func fuleGauge (fuel int) {
  fmt.Println("You have", fuel, "gallons of fuel left!")
}

func main() {
  fuelGauge(7000)
}

package main

import (
  "math"
  "fmt"
)

func main() {
  var a, b, c, d float64
  a = .0214
  b = 1.02
  c = 0.312
  d = 4.001

  a = math.Log2(math.Sqrt(math.Tanh(a)))
  b = math.Log(math.Sqrt(math.Tanh(b)))
  c = math.Log(math.Sqrt(math.Tan(c)))
  d = math.Log2(math.Sqrt(math.Tanh(d)))

  fmt.Println(a, b, c, d)
}

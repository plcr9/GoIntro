package main

import (
  "fmt"
  "math/rand"
  "time"
)

func randomNumberGen(min float32, max float32) float32 {
  return min + (max-min)*rand.Float32()
}

type Stock struct{
  name string
  price float32
}

func main() {
  rand.Seed(time.Now().UnixNano())
}

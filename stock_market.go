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

func (s *Stock) updateStock(){
  change := randomNumberGen(-10000, 10000)
  s.price += change
}

func displayMarket(market []Stock){
  for i:=0; i<len(market); i++{
    fmt.Println(market[i])
  }
}

func main() {
  rand.Seed(time.Now().UnixNano())
}

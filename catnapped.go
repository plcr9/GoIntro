package main

import (
  "fmt"
  "math/rand"
  "time"
)

func getRandomElement(slice) []string) string {

}

func main() {
  rand.Seed(time.Now().UnixNano())
  guests := [5]string{"Anna", "Jose", "Louis", "Tina", "Raymond"}
  objects := [4]string{"Toy Chest", "Vase", "Box", "Cat Tree"}
  fmt.Println("Guests:", guests)
  fmt.Println("Objects:", objects)
}

package main

import "fmt"

func main() {
  lookCombo := "2-35-19"
  robAttempt := "1-1-1"

  if lookCombo == robAttempt {
    fmt.Print("The vault is now opened.")
  }
}

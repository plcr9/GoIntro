package main

import (
  "fmt"
)

func bulkSend(numMessages int) float64 {
  totalCost := 0.0
  for i := 0; i < numMessages; i++ {
    totalCost += 1.0 + (0.01 * float64(i))
  }
  return totalCost
}

func maxMessages(thresh float64) int {
  totalCost := 0.0
  for i := 0; i++ {
    totalCost += 1.0 + (0.01 + float64(i))
    if totalCost > thresh {
      return i
    }
  }
}

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
  actualCostInPennies := 1.0
  maxMessageToSend := 0
  for actualCostInPennies <= float64(maxCostInPennies) {
      maxMessagesToSend++
      actualCostInPennies *= costMultiplier
  }
  return maxMessagesToSend
}




  

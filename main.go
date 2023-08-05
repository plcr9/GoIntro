package main

import "fmt"

func main () {
  var currentStation string
  var nextTrainTime int8
  var isUptownTrain bool

  currentStation = "Aldgate East"
  nextTrainTime = 4
  isUptownTrain = true

  fmt.Println("Current station:", currentStation)
  fmt.Println("Next train:", nextTrainTime)
  fmt.Println("Is uptown:", isUptownTrain)

  currentStation = "Grand Central"
  nextTrainTime = 3
  isUptownTrain = true

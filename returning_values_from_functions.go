package main

import (
  "fmt"
  "time"
)

func isItLateInNewYork() string {
  var lateMessage string
  t := time.Now()
  tz, _ := time.LoadLocation("America/New_York")
  myHour := t.In(tz).Hour()
  if myHour < 5 {
    lateMessage = "Goodness it is late"
  } else if myHour < 16 {
    lateMessage = "It's not late at all!"
  } else if myHour < 19 {
    lateMessage = "I guess it's getting kind of late"
  } else {
    lateMessage = "It's late"
  }

  return lateMessage
}

func main() {
  var myLate string
  nyLate = isItLateInNewYork()
  fmt.Println(nyLate)
}

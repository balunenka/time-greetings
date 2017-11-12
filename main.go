package main

import (
	"fmt"
	"time"
)

// Returns greeting text according to get current hour
// It need a parameter int hour - 0...24
func getGreetings(hour int) string {

	switch {
	case (hour >= 6) && (hour < 12):
		return "Good morning!"
	case (hour >= 12) && (hour < 18):
		return "Good afternoon!"
	case (hour >= 18) && (hour < 23):
		return "Good evening!"
	case (hour >= 23) && (hour < 6):
		return "Good night!"
	default:
		return "Incorrect time......"
	}

}

/*
	TODO


*/
func main() {
	hourOfDay := time.Now().Hour()
	minutesOfDay := time.Now().Minute()
	todayDay := time.Now().Day()
	weekdayOfDay := time.Now().Weekday()

	fmt.Println("It's", hourOfDay, ":", minutesOfDay, " now")
	fmt.Println("Today is ", weekdayOfDay, todayDay)

	fmt.Println(getGreetings(hourOfDay))
	//fmt.Println(getGreetings(1)) //  for test

}

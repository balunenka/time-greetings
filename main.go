package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

// Returns greeting text according to get current hour
// It need a parameter int hour - 0...24
func getGreetings(hour int) (string, error) {
	var message string
	switch {
	case (hour >= 6) && (hour < 12):
		message = "Good morning!"
	case (hour >= 12) && (hour < 18):
		message = "Good afternoon!"
	case (hour >= 18) && (hour < 22):
		message = "Good evening!"
	case (((hour >= 22) && (hour <= 24)) || ((hour >= 0) && (hour < 6))):
		message = "Good night!"
	default:
		err := errors.New("[ERROR] Incorrect value for hour: " + strconv.Itoa(hour))
		return message, err
	}

	return message, nil
}

/*
	TODO


*/
func main() {
	hourOfDay := time.Now().Hour()
	minutesOfDay := time.Now().Minute()
	//timeNow := strconv.Itoa( hourOfDay )+":"+strconv.Itoa( minutesOfDay )

	todayDay := time.Now().Day()
	weekdayOfDay := time.Now().Weekday().String()
	monthNow := time.Now().UTC().Format("January")
	//dateNowMessage := "Today is: " + weekdayOfDay[ 0:3 ] + ", " + strconv.Itoa( todayDay ) + " of " + monthNow

	fmt.Printf("It's %d:%d now \n", hourOfDay, minutesOfDay)
	fmt.Printf("Today is: %s, %d of %s \n", weekdayOfDay, todayDay, monthNow)

	greeting, err := getGreetings(hourOfDay)

	if err != nil {
		color.Red(err.Error()) // print error message "err" in color red
		os.Exit(1)

	}
	fmt.Println(greeting)

}

package main

import (
	"errors"
	"flag"
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

// User input of hour
// returns error if not int entered and looped while not succeed
func setHour() int {

	var hour int
	var isEntered bool
	for !isEntered {
		fmt.Printf("Please enter the hour: ")
		_, err := fmt.Scan(&hour)

		if err != nil {
			color.Red(err.Error()) // print error message "err" in color red if not integer was inputted

		} else {
			isEntered = true
		}

	}

	fmt.Printf("\n\n\n")
	return hour
}

// Get data and time values
func getCurrentTime() (int, int, int, string, string) {

	hour := time.Now().Hour()
	minutes := time.Now().Minute()

	day := time.Now().Day()
	weekday := time.Now().Weekday().String()
	month := time.Now().UTC().Format("January")

	return hour, minutes, day, weekday, month
}

// Main function that runs the program
func main() {
	var isTestMode bool
	flag.BoolVar(&isTestMode, "testMode", false, "Turns on test mode of application. Ability to input manually hour")
	flag.Parse()

	var hourOfDay, minutesOfDay, todayDay int
	var weekdayOfDay, monthNow string

	todayDay = -1     // set as default incorrect value
	minutesOfDay = -1 // set as default incorrect value
	hourOfDay = -1    // set as default incorrect value

	if isTestMode {
		fmt.Printf("In a test mode now!\n\n")
		_, minutesOfDay, todayDay, weekdayOfDay, monthNow = getCurrentTime()
		hourOfDay = setHour()

		fmt.Printf("It's %d:%d now \n", hourOfDay, minutesOfDay)
		fmt.Printf("Today is: %s, %d of %s \n\n", weekdayOfDay, todayDay, monthNow)

	} else {
		fmt.Printf("In ususal mode...\n\n")

		hourOfDay, minutesOfDay, todayDay, weekdayOfDay, monthNow = getCurrentTime()

		fmt.Printf("It's %d:%d now \n", hourOfDay, minutesOfDay)
		fmt.Printf("Today is: %s, %d of %s \n\n", weekdayOfDay, todayDay, monthNow)
	}

	greeting, err := getGreetings(hourOfDay)

	if err != nil {
		color.Red(err.Error()) // print error message "err" in color red
		os.Exit(1)

	}
	fmt.Println(greeting)

}

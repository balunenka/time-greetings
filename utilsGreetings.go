package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Returns greeting text according to get current hour
// It need a parameter int hour - 0...24
func getGreetings(hour int) (string, error) {

	var message string
	var err error

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
		//err := errors.New("[ERROR] Incorrect value for hour: " + strconv.Itoa(hour))
		err = greetingError(strconv.Itoa(hour))

	}

	return message, err
}

// User input of hour
// returns error if not int entered and looped while not succeed
func setHour() (bool, int, error) {

	var (
		hour      int    // hour that is get from converting user input to integer
		input     string // user input
		errReader error  // error for reader.ReadString - input
		errAtoi   error  // error for strconv.Atoi - converting of user string input to integer
		err       error  // error message that will be returned from function setHour

		isEntered bool // indicate if input was successful
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter text: ")
	input, errReader = reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	if !(checkError(errReader)) {
		hour, errAtoi = strconv.Atoi(input) // convert user string input to int
		if !(checkError(errAtoi)) {
			isEntered = true // check if conversation was correct - exit loop of input

		} else {
			err = inputError(input)
		}

	}

	fmt.Printf("\n\n\n")
	return isEntered, hour, err
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

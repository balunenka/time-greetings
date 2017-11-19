package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	// ErrInput indicates that a value is out of range for the target type.
	ErrInput = errors.New("should be integer in range [0;24]")

	// ErrGreeting indicates that a value does not have the right syntax for the target type.
	ErrGreeting = errors.New("hour out of range [0;24]")
)

// GretingsError records for what input error ocurs and with what value
type GreetingsError struct {
	Err       error  // the reason why failed (ErrInput, ErrGreeting)
	HourInput string // the user input
}

// function constructing error message from  GreetingsError struct
func (e *GreetingsError) Error() string {
	return "[ERROR] " + " '" + e.HourInput + "' " + e.Err.Error()
}

// function sets structure GreetingsError
func inputError(str string) *GreetingsError {
	return &GreetingsError{
		ErrInput,
		str,
	}
}

// function sets structure GreetingsError
func greetingError(str string) *GreetingsError {
	return &GreetingsError{
		ErrGreeting,
		str,
	}
}

// error catcher - returns true value
func checkError(err error) bool {
	var errorExist bool
	if err != nil {
		errorExist = true
		color.Red(err.Error())
		log.Printf(err.Error())

	} else {
		errorExist = false
	}
	return errorExist
}

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

// Main function that runs the program
func main() {
	var isTestMode bool
	flag.BoolVar(&isTestMode, "testMode", false, "Turns on test mode of application. Ability to input manually hour")
	flag.Parse()

	var hourOfDay, minutesOfDay, todayDay int
	var weekdayOfDay, monthNow string
	var isEntered bool
	var setHourErr, greetingErr error

	todayDay = -1     // set as default incorrect value
	minutesOfDay = -1 // set as default incorrect value
	hourOfDay = -1    // set as default incorrect value

	if isTestMode {
		color.HiGreen("Running in a test mode now!\n\n")
		_, minutesOfDay, todayDay, weekdayOfDay, monthNow = getCurrentTime()

		for {
			isEntered, hourOfDay, setHourErr = setHour()
			if isEntered {
				break
			} else {
				if checkError(setHourErr) {
					color.Yellow("Try again...\n\n")
				}

			}
		}

	} else {
		color.HiGreen("Running in an ususal mode...\n\n")

		hourOfDay, minutesOfDay, todayDay, weekdayOfDay, monthNow = getCurrentTime()
	}

	greeting, greetingErr := getGreetings(hourOfDay)

	if !(checkError(greetingErr)) {
		fmt.Printf("It's %d:%d now \n", hourOfDay, minutesOfDay)
		fmt.Printf("Today is: %s, %d of %s \n\n", weekdayOfDay, todayDay, monthNow)

		fmt.Println(greeting)
	} else {
		os.Exit(1)
	}

}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

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

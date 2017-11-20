package main

import (
	"errors"
	"log"

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

package main

import (
	"errors"
	"strconv"
	"strings"
)

// KeyAction represents every line we get from hid_listen,
// those lines describe a key press on the keyboard
type KeyAction struct {
	Keyboard       string
	Column         int
	Row            int
	Press          bool
	TapCount       int
	TapInterrupted bool
	Keycode        int
	Layer          int
}

// stringConverter is used to convert string values
// into the desired types. We use this techniques to make error handling easier
// when you need to convert multiple values in a row.
// By holding the error value in a object we can make multiple convertions
// and only check for errors at the end.
type stringConverter struct {
	err error
}

func (sc *stringConverter) toInt(text string) int {
	result := 0

	if sc.err != nil {
		return result
	}

	result, sc.err = strconv.Atoi(text)

	return result
}

func (sc *stringConverter) toBool(text string) bool {
	result := false

	if sc.err != nil {
		return result
	}

	result, sc.err = strconv.ParseBool(text)

	return result
}

// NewKeyAction serializes the input string into a KeyAction and returns it
func NewKeyAction(input string) (*KeyAction, error) {
	splitInput := strings.Split(input, ",")
	if len(splitInput) != 8 {
		return nil, errors.New("Invalid input. It should contain 7 fields")
	}

	converter := &stringConverter{}
	keyAction := &KeyAction{}

	keyAction.Keyboard = splitInput[0]
	keyAction.Column = converter.toInt(splitInput[1])
	keyAction.Row = converter.toInt(splitInput[2])
	keyAction.Press = converter.toBool(splitInput[3])
	keyAction.TapCount = converter.toInt(splitInput[4])
	keyAction.TapInterrupted = converter.toBool(splitInput[5])
	keyAction.Keycode = converter.toInt(splitInput[6])
	keyAction.Layer = converter.toInt(splitInput[7])

	if converter.err != nil {
		return nil, converter.err
	}

	return keyAction, nil
}

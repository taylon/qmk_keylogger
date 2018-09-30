package main

import (
	"errors"
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

// NewKeyAction serializes the input string into a KeyAction and returns it
func NewKeyAction(input string) (*KeyAction, error) {
	splitInput := strings.Split(input, ",")
	if len(splitInput) != 8 {
		return nil, errors.New("Invalid input. It should contain 7 fields")
	}

	converter := &StringConverter{}
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

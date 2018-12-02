package keyaction

import (
	"fmt"
	"strings"

	"github.com/taylon/qmk_keylogger/utils"
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
	KeyCode        int
	Layer          int
}

// New serializes the input string into a KeyAction and returns it
func New(input string) (*KeyAction, error) {
	splitInput := strings.Split(input, ",")
	if len(splitInput) != 8 {
		return nil, fmt.Errorf("invalid input. It should contain 7 fields: %s", input)
	}

	converter := &utils.StringConverter{}
	keyAction := &KeyAction{}

	keyAction.Keyboard = splitInput[0]
	keyAction.Column = converter.ToInt(splitInput[1])
	keyAction.Row = converter.ToInt(splitInput[2])
	keyAction.Press = converter.ToBool(splitInput[3])
	keyAction.TapCount = converter.ToInt(splitInput[4])
	keyAction.TapInterrupted = converter.ToBool(splitInput[5])
	keyAction.KeyCode = converter.ToInt(splitInput[6])
	keyAction.Layer = converter.ToInt(splitInput[7])

	if converter.Error != nil {
		return nil, converter.Error
	}

	return keyAction, nil
}

package keyaction

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKeyPress(t *testing.T) {
	expectedKeyAction := &KeyAction{
		Keyboard:       "dactyl",
		Column:         11,
		Row:            2,
		Press:          true,
		TapCount:       0,
		TapInterrupted: false,
		KeyCode:        26,
		Layer:          1,
	}

	KeyAction, _ := New(fmt.Sprintf("%s,%d,%d,%t,%d,%t,%d,%d",
		expectedKeyAction.Keyboard,
		expectedKeyAction.Column,
		expectedKeyAction.Row,
		expectedKeyAction.Press,
		expectedKeyAction.TapCount,
		expectedKeyAction.TapInterrupted,
		expectedKeyAction.KeyCode,
		expectedKeyAction.Layer,
	))

	assert.Equal(t, expectedKeyAction, KeyAction)
}

func TestNewKeyPressReturnsErrorWhenInvalidInput(t *testing.T) {
	// Test an input that does not have at least 8 fields
	_, err := New("mx_browns_sucks")
	assert.Error(t, err)

	// Test an input that has 8 fields but is still invalid
	_, err = New("zealios62!,zealios62!,zealios62!,zealios62!,zealios62!,zealios62!,zealios62!,zealios62!")
	assert.Error(t, err)

	// Test an input that has 7 valid fields and one invalid
	_, err = New("invalid,invalid,invalid,0,invalid,invalid,invalid,invalid")
	assert.Error(t, err)
}

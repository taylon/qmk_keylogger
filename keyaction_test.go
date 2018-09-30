package main

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
		Keycode:        26,
		Layer:          1,
	}

	KeyAction, _ := NewKeyAction(fmt.Sprintf("%s,%d,%d,%t,%d,%t,%d,%d",
		expectedKeyAction.Keyboard,
		expectedKeyAction.Column,
		expectedKeyAction.Row,
		expectedKeyAction.Press,
		expectedKeyAction.TapCount,
		expectedKeyAction.TapInterrupted,
		expectedKeyAction.Keycode,
		expectedKeyAction.Layer,
	))

	assert.Equal(t, expectedKeyAction, KeyAction)
}

func TestNewKeyPressReturnsErrorWhenInvalidInput(t *testing.T) {
	// Test an input that does not have at least 8 fields
	_, err := NewKeyAction("mx_browns_sucks")
	assert.Error(t, err)

	// Test an input that has 8 fields but is still invalid
	_, err = NewKeyAction("zealios62!,zealios62!,zealios62!,zealios62!,zealios62!,zealios62!,zealios62!,zealios62!")
	assert.Error(t, err)

	// Test an input that has 7 valid fields and one invalid
	_, err = NewKeyAction("invalid,invalid,invalid,0,invalid,invalid,invalid,invalid")
	assert.Error(t, err)
}

func TestStringConverterToInt(t *testing.T) {
	// Sucessfull call
	converter := &stringConverter{}
	converter.toInt("1")
	converter.toInt("2")
	result := converter.toInt("0")
	assert.Equal(t, 0, result)
	assert.NoError(t, converter.err)

	// Call with invalid input
	converter = &stringConverter{}
	converter.toInt("mx_browns_sucks")
	assert.Error(t, converter.err)

	// Call with two invalid calls and one valid
	converter = &stringConverter{}
	converter.toInt("0")
	converter.toInt("mx_browns_sucks")
	converter.toInt("0")
	assert.Error(t, converter.err)
}

func TestStringConverterToBool(t *testing.T) {
	// Sucessfull call
	converter := &stringConverter{}
	converter.toBool("0")
	resultFalse := converter.toBool("0")
	resultTrue := converter.toBool("1")
	assert.False(t, resultFalse)
	assert.True(t, resultTrue)
	assert.NoError(t, converter.err)

	// Call with invalid input
	converter = &stringConverter{}
	converter.toBool("mx_browns_sucks")
	assert.Error(t, converter.err)

	// Call with two invalid calls and one valid
	converter = &stringConverter{}
	converter.toBool("0")
	converter.toBool("mx_browns_sucks")
	converter.toBool("2")
	assert.Error(t, converter.err)
}

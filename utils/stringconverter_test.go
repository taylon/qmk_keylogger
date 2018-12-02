package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringConverterToInt(t *testing.T) {
	// Successful call
	converter := &StringConverter{}
	converter.ToInt("1")
	converter.ToInt("2")

	assert.Equal(t, 0, converter.ToInt("0"))
	assert.NoError(t, converter.Error)

	// Call with invalid input
	converter = &StringConverter{}
	converter.ToInt("mx_browns_sucks")

	assert.Error(t, converter.Error)

	// Call with two invalid calls and one valid
	converter = &StringConverter{}
	converter.ToInt("0")
	converter.ToInt("mx_browns_sucks")
	converter.ToInt("0")

	assert.Error(t, converter.Error)
}

func TestStringConverterToBool(t *testing.T) {
	// Successful call
	converter := &StringConverter{}
	converter.ToBool("0")

	assert.False(t, converter.ToBool("0"))
	assert.True(t, converter.ToBool("1"))
	assert.NoError(t, converter.Error)

	// Call with invalid input
	converter = &StringConverter{}
	converter.ToBool("mx_browns_sucks")

	assert.Error(t, converter.Error)

	// Call with two invalid calls and one valid
	converter = &StringConverter{}
	converter.ToBool("0")
	converter.ToBool("mx_browns_sucks")
	converter.ToBool("2")

	assert.Error(t, converter.Error)
}

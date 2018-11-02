package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringConverterToInt(t *testing.T) {
	// Successful call
	converter := &StringConverter{}
	converter.toInt("1")
	converter.toInt("2")

	assert.Equal(t, 0, converter.toInt("0"))
	assert.NoError(t, converter.err)

	// Call with invalid input
	converter = &StringConverter{}
	converter.toInt("mx_browns_sucks")

	assert.Error(t, converter.err)

	// Call with two invalid calls and one valid
	converter = &StringConverter{}
	converter.toInt("0")
	converter.toInt("mx_browns_sucks")
	converter.toInt("0")

	assert.Error(t, converter.err)
}

func TestStringConverterToBool(t *testing.T) {
	// Successful call
	converter := &StringConverter{}
	converter.toBool("0")

	assert.False(t, converter.toBool("0"))
	assert.True(t, converter.toBool("1"))
	assert.NoError(t, converter.err)

	// Call with invalid input
	converter = &StringConverter{}
	converter.toBool("mx_browns_sucks")

	assert.Error(t, converter.err)

	// Call with two invalid calls and one valid
	converter = &StringConverter{}
	converter.toBool("0")
	converter.toBool("mx_browns_sucks")
	converter.toBool("2")

	assert.Error(t, converter.err)
}

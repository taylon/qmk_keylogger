package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringConverterToInt(t *testing.T) {
	// Sucessfull call
	converter := &StringConverter{}
	converter.toInt("1")
	converter.toInt("2")
	result := converter.toInt("0")
	assert.Equal(t, 0, result)
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
	// Sucessfull call
	converter := &StringConverter{}
	converter.toBool("0")
	resultFalse := converter.toBool("0")
	resultTrue := converter.toBool("1")
	assert.False(t, resultFalse)
	assert.True(t, resultTrue)
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

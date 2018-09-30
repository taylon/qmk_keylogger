package main

import "strconv"

// StringConverter is used to convert string values
// into the desired types. We use this technique to make error handling easier
// when you need to convert multiple values in a row.
// By holding the error value in a object we can make multiple convertions
// and only check for errors at the end.
type StringConverter struct {
	err error
}

func (sc *StringConverter) toInt(text string) int {
	result := 0

	if sc.err != nil {
		return result
	}

	result, sc.err = strconv.Atoi(text)

	return result
}

func (sc *StringConverter) toBool(text string) bool {
	result := false

	if sc.err != nil {
		return result
	}

	result, sc.err = strconv.ParseBool(text)

	return result
}

package utils

import "strconv"

// StringConverter is used to convert string values
// into the desired types. We use this technique to make error handling easier
// when you need to convert multiple values in a row.
// By holding the error value in a object we can make multiple conversions
// and only check for errors at the end.
type StringConverter struct {
	Error error
}

func (sc *StringConverter) ToInt(text string) int {
	if sc.Error != nil {
		return 0
	}

	var result int
	result, sc.Error = strconv.Atoi(text)

	return result
}

func (sc *StringConverter) ToBool(text string) bool {
	if sc.Error != nil {
		return false
	}

	var result bool
	result, sc.Error = strconv.ParseBool(text)

	return result
}

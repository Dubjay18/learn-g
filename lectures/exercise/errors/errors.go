//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hours, minute, second int
}

type TimeParserError struct {
	message string
	input   string
}

func (t *TimeParserError) Error() string {
	return fmt.Sprintf("%v : %v", t.message, t.input)

}

func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &TimeParserError{"invalid time string", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParserError{fmt.Sprintf("invalid hour: %v", err), input}
		}
		minute, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimeParserError{fmt.Sprintf("invalid minute: %v", err), input}
		}
		second, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &TimeParserError{fmt.Sprintf("invalid second: %v", err), input}
		}
		if hour < 0 || hour > 23 {
			return Time{}, &TimeParserError{"hour out of range", input}
		}
		if minute < 0 || minute > 59 {
			return Time{}, &TimeParserError{"minute out of range", input}
		}
		if second < 0 || second > 59 {
			return Time{}, &TimeParserError{"second out of range", input}
		}

		return Time{hour, minute, second}, nil
	}
}

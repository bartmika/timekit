package timekit

import (
	"time"
)

// YearsRange returns an array of the year integer values between two dates. For example if one date is 2000 and the other is 2002, the output will be [2000,2001,2002].
func YearsRange(start time.Time, end time.Time) []int {
	var years []int

	// Developers Note:
	// We want to leverage our already unit tested code for the `range` functionality so we will use the `TimeStepper`
	// to iterate through the datetime values and add them to an `results` array.
	ts := NewTimeStepper(start, end, 1, 0, 0, 0, 0, 0)
	running := true
	for running {
		// Get the value we are on in the timestepper.
		v := ts.Value()

		years = append(years, v.Year())

		// Run our timestepper to get our next value.
		ts.Step()

		running = ts.Done() == false
	}
	return years
}

// MonthRange returns an array of the month integer values between two dates. For example if one date is January 2000 and the other is March 2000, the output will be [1,2,3].
func MonthRange(start time.Time, end time.Time) []int {
	var months []int

	// Developers Note:
	// We want to leverage our already unit tested code for the `range` functionality so we will use the `TimeStepper`
	// to iterate through the datetime values and add them to an `results` array.
	ts := NewTimeStepper(start, end, 0, 1, 0, 0, 0, 0)
	running := true
	for running {
		// Get the value we are on in the timestepper.
		v := ts.Value()

		// Developers Note:
		// (1) Yes the 'time.Month' type is a string as per https://pkg.go.dev/time#Month
		// (2) But the underlying type is int, so it can be converted to int. (https://stackoverflow.com/a/16686866)
		// (3) As a result this code will work.
		months = append(months, int(v.Month()))

		// Run our timestepper to get our next value.
		ts.Step()

		running = ts.Done() == false
	}
	return months
}

// DaysRange returns an array of the day integer values between two dates. For example if one date is January 1st 2000 and the other is January 5th 2000, the output will be [1,2,3,4,5].
func DaysRange(start time.Time, end time.Time) []int {
	var days []int

	// Developers Note:
	// We want to leverage our already unit tested code for the `range` functionality so we will use the `TimeStepper`
	// to iterate through the datetime values and add them to an `results` array.
	ts := NewTimeStepper(start, end, 0, 0, 1, 0, 0, 0)
	running := true
	for running {
		// Get the value we are on in the timestepper.
		v := ts.Value()

		days = append(days, v.Day())

		// Run our timestepper to get our next value.
		ts.Step()

		running = ts.Done() == false
	}
	return days
}

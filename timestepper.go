package timekit

import (
	"time"
)

// TimeStepper is a structure to hold keep track of the position we are in the datetime range which we are stepping through.
type TimeStepper struct {
	tz         *time.Location
	curr       time.Time
	start      time.Time
	end        time.Time
	yearStep   int
	monthStep  int
	dayStep    int
	hourStep   int
	minuteStep int
	secondStep int
}

// NewTimeStepper is a constructor of the `TimeStepper` struct.
func NewTimeStepper(start time.Time, end time.Time, yearStep int, monthStep int, dayStep int, hourStep int, minuteStep int, secondStep int) *TimeStepper {
	return &TimeStepper{
		tz:         start.Location(),
		curr:       start,
		start:      start,
		end:        end,
		yearStep:   yearStep,
		monthStep:  monthStep,
		dayStep:    dayStep,
		hourStep:   hourStep,
		minuteStep: minuteStep,
		secondStep: secondStep,
	}
}

// Step makes one time step over and returns true or false depending if the stepper has stepped over the end datetime.
func (ts *TimeStepper) Step() bool {
	// Developers Note:
	// It was discovered that when you run the stepper in situtations involving
	// daylight saving time, the stepper would lock up in a single date and do
	// not proceed. This was found when you want to run the timestepper from
	// Jan 1th 2021 to Jan 1th 2022 in 5 minute step intervals under the
	// "America/Toronto" timezone. In essence daylight savings are the problem.
	// To find a solution, stackoverflow had a great answer in the following:
	// https://stackoverflow.com/questions/5495803/does-utc-observe-daylight-saving-time
	// The `TL;DR;` is UTC does not have daylight savings, so therefore I am
	// converting the current time into a UTC timezone value, performing our step
	// and then converting back to the specific local timezone - therfore fixing
	// our issue.
	currUTC := ts.curr.UTC()
	dur := time.Hour*time.Duration(ts.hourStep) + time.Minute*time.Duration(ts.minuteStep) + time.Second*time.Duration(ts.secondStep)
	currUTC = currUTC.AddDate(ts.yearStep, ts.monthStep, ts.dayStep).Add(dur)
	ts.curr = currUTC.In(ts.tz)

	// Returns true or false depending if the stepper has stepped over the end datetime.
	return ts.curr.After(ts.end) == false
}

// Done checks to see if the stepper has stepped over the end datetime and will return true or false according.
func (ts *TimeStepper) Done() bool {
	return ts.curr.After(ts.end)
}

// Value will return the value that that the stepper is currently on.
func (ts *TimeStepper) Value() time.Time {
	return ts.curr
}

// RangeFromTimeStepper function returns an array of datetime values from the starting date up to and including the finish date according to the step pattern specified in the parameter.
func RangeFromTimeStepper(start time.Time, end time.Time, yearStep int, monthStep int, dayStep int, hourStep int, minuteStep int, secondStep int) []time.Time {
	var results []time.Time

	// Developers Note:
	// We want to leverage our already unit tested code for the `range` functionality so we will use the `TimeStepper`
	// to iterate through the datetime values and add them to an `results` array.
	ts := NewTimeStepper(start, end, yearStep, monthStep, dayStep, hourStep, minuteStep, secondStep)
	running := true
	for running {
		// Get the value we are on in the timestepper.
		v := ts.Value()

		results = append(results, v)

		// Run our timestepper to get our next value.
		ts.Step()

		running = ts.Done() == false
	}
	return results
}

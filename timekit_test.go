package timekit

import (
	"reflect"
	"testing"
	"time"
)

// Developers Note:
// Special thanks to the following article which helped explain how to stub
// out the `time.Now()` function in Golang.
// https://labs.yulrizka.com/en/stubbing-time-dot-now-in-golang/

func TestFirstDayOfLastYear(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 1, 1, 0, 0, 0, 0, loc) // Jan 1st 2000
	}

	actual := FirstDayOfLastYear(timeFn)
	expected := time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC) // Jan 1st 1999
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestFirstDayOfThisYear(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 7, 1, 0, 0, 0, 0, loc) // July 1st 2000
	}

	actual := FirstDayOfThisYear(timeFn)
	expected := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC) // Jan 1st 2000
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestFirstDayOfNextYear(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2022, 7, 1, 0, 0, 0, 0, loc) // July 1st 2022
	}

	actual := FirstDayOfNextYear(timeFn)
	expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC) // Jan 1st 2023
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestFirstDayOfLastMonth(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 1, 1, 0, 0, 0, 0, loc) // Jan 1st 2000
	}

	actual := FirstDayOfLastMonth(timeFn)
	expected := time.Date(1999, 12, 1, 0, 0, 0, 0, time.UTC) // Dec 1st 1999
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestFirstDayOfThisMonth(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 1, 20, 0, 0, 0, 0, loc) // Jan 20th 2000
	}

	actual := FirstDayOfThisMonth(timeFn)
	expected := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC) // Jan 1st 2000
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestFirstDayOfNextMonth(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 1, 20, 0, 0, 0, 0, loc) // Jan 20th 2000
	}

	actual := FirstDayOfNextMonth(timeFn)
	expected := time.Date(2000, 2, 1, 0, 0, 0, 0, time.UTC) // Feb 1st 2000
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestMidnightYesterday(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 1, 20, 0, 0, 0, 0, loc) // Jan 20th 2000 12AM
	}

	actual := MidnightYesterday(timeFn)
	expected := time.Date(2000, 1, 19, 0, 0, 0, 0, time.UTC) // Jan 19th 2000 12AM
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestMidnight(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 1, 20, 5, 0, 0, 0, loc) // Jan 20th 2000 5AM
	}

	actual := Midnight(timeFn)
	expected := time.Date(2000, 1, 20, 0, 0, 0, 0, time.UTC) // Jan 20th 2000 12AM
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestMidnightTomorrow(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 1, 20, 5, 0, 0, 0, loc) // Jan 20th 2000 5AM
	}

	actual := MidnightTomorrow(timeFn)
	expected := time.Date(2000, 1, 21, 0, 0, 0, 0, time.UTC) // Jan 21st 2000 12AM
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestNoon(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 1, 20, 5, 0, 0, 0, loc) // Jan 20th 2000 5AM
	}

	actual := Noon(timeFn)
	expected := time.Date(2000, 1, 20, 12, 0, 0, 0, time.UTC) // Jan 200th 2000 12PM
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestFirstDayOfLastISOWeek(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2022, 1, 9, 1, 0, 0, 0, loc) // Sunday Jan 9th
	}

	actual := FirstDayOfLastISOWeek(timeFn)
	expected := time.Date(2021, 12, 27, 1, 0, 0, 0, time.UTC) // Monday Jan 3ed
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestFirstDayOfThisISOWeek(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2022, 1, 9, 1, 0, 0, 0, loc) // Sunday Jan 9th
	}

	actual := FirstDayOfThisISOWeek(timeFn)
	expected := time.Date(2022, 1, 3, 1, 0, 0, 0, time.UTC) // Monday Jan 3ed
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestLastDayOfThisISOWeek(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2022, 1, 7, 1, 0, 0, 0, loc) // Friday Jan 7th
	}

	actual := LastDayOfThisISOWeek(timeFn)
	expected := time.Date(2022, 1, 9, 1, 0, 0, 0, time.UTC) // Sunday Jan 9th
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestFirstDayOfNextISOWeek(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2022, 1, 7, 1, 0, 0, 0, loc) // Friday Jan 7th
	}

	actual := FirstDayOfNextISOWeek(timeFn)
	expected := time.Date(2022, 1, 10, 1, 0, 0, 0, time.UTC) // Monday Jan 10th
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestRange(t *testing.T) {
	loc := time.UTC                                 // closure can be used if necessary
	start := time.Date(2022, 1, 7, 1, 0, 0, 0, loc) // Jan 7th 2022
	end := time.Date(2022, 1, 10, 1, 0, 0, 0, loc)  // Jan 10th 2022
	actual := Range(start, end, 0, 0, 1, 0, 0, 0)
	expected := []time.Time{
		time.Date(2022, 1, 7, 1, 0, 0, 0, loc),  // Jan 7th 2022
		time.Date(2022, 1, 8, 1, 0, 0, 0, loc),  // Jan 8th 2022
		time.Date(2022, 1, 9, 1, 0, 0, 0, loc),  // Jan 9th 2022
		time.Date(2022, 1, 10, 1, 0, 0, 0, loc), // Jan 10th 2022
	}
	if reflect.DeepEqual(actual, expected) == false {
		t.Errorf("Incorrect date ranges, got %s but was expecting %s", actual, expected)
	}
}

func TestNewTimeStepper(t *testing.T) {
	loc := time.UTC                                 // closure can be used if necessary
	start := time.Date(2022, 1, 7, 1, 0, 0, 0, loc) // Jan 7th 2022
	end := time.Date(2022, 1, 10, 1, 0, 0, 0, loc)  // Jan 10th 2022
	actual := NewTimeStepper(start, end, 0, 0, 1, 0, 0, 0)
	expected := &TimeStepper{
		tz:         loc,
		curr:       start,
		start:      start,
		end:        end,
		yearStep:   0,
		monthStep:  0,
		dayStep:    1,
		hourStep:   0,
		minuteStep: 0,
		secondStep: 0,
	}
	if reflect.DeepEqual(actual, expected) == false {
		t.Errorf("Incorrect struct, got %v but was expecting %v", actual, expected)
	}
}

func TestStep(t *testing.T) {
	loc := time.UTC                                 // closure can be used if necessary
	start := time.Date(2022, 1, 7, 1, 0, 0, 0, loc) // Jan 7th 2022
	end := time.Date(2022, 1, 10, 1, 0, 0, 0, loc)  // Jan 10th 2022
	ts := NewTimeStepper(start, end, 0, 0, 1, 0, 0, 0)

	// Perform our step operation and verify it works.
	isFinished := ts.Step()

	if isFinished == false {
		t.Errorf("Incorrect value, got %v but was expecting %v", isFinished, true)
	}
}

func TestDone(t *testing.T) {
	loc := time.UTC                                 // closure can be used if necessary
	start := time.Date(2022, 1, 7, 1, 0, 0, 0, loc) // Jan 7th 2022
	end := time.Date(2022, 1, 10, 1, 0, 0, 0, loc)  // Jan 10th 2022
	ts := NewTimeStepper(start, end, 0, 0, 1, 0, 0, 0)

	// Perform our step operation and verify it works.
	isFinished := ts.Done()

	if isFinished == true {
		t.Errorf("Incorrect value, got %v but was expecting %v", isFinished, false)
	}
}

func TestValue(t *testing.T) {
	loc := time.UTC                                 // closure can be used if necessary
	start := time.Date(2022, 1, 7, 1, 0, 0, 0, loc) // Jan 7th 2022
	end := time.Date(2022, 1, 10, 1, 0, 0, 0, loc)  // Jan 10th 2022
	ts := NewTimeStepper(start, end, 0, 0, 1, 0, 0, 0)

	// Without step operation.
	expected := time.Date(2022, 1, 7, 1, 0, 0, 0, loc) // Jan 7th 2022
	actual := ts.Value()
	if expected != actual {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}

	// Perform our step operation and let's verify the value is correct.
	ts.Step()
	expected = time.Date(2022, 1, 8, 1, 0, 0, 0, loc) // Jan 8th 2022
	actual = ts.Value()
	if expected != actual {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}

	// Perform yet another step and verify value function.
	ts.Step()
	expected = time.Date(2022, 1, 9, 1, 0, 0, 0, loc) // Jan 9th 2022
	actual = ts.Value()
	if expected != actual {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

//TODO: IMPL.
func TestTimeStepperTorontoBug(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")
	start := time.Date(2021, 1, 1, 1, 0, 0, 0, loc) // Jan 1th 2021
	end := time.Date(2022, 1, 1, 1, 0, 0, 0, loc)   // Jan 1th 2022
	ts := NewTimeStepper(start, end, 0, 0, 0, 0, 5, 0)

	var actual time.Time
	running := true
	for running {
		// Get the value we are on in the timestepper.
		actual = ts.Value()

		// log.Println(actual) // For debugging purposes only.

		// Run our timestepper to get our next value.
		ts.Step()

		running = ts.Done() == false
	}

	expected := time.Date(2022, 1, 1, 1, 0, 0, 0, loc) // Jan 1th 2022
	if expected != actual {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

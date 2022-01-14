package timekit

import (
	"reflect"
	"testing"
	"time"
)

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

// Developers note: This is to verify the non-UTC timezone daylight saving issue is fixed.
func TestTimeStepperTorontoTimezoneBug(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")
	start := time.Date(2021, 7, 1, 1, 0, 0, 0, loc) // July 1st 2021
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

func TestRangeFromTimeStepper(t *testing.T) {
	loc := time.UTC                                 // closure can be used if necessary
	start := time.Date(2022, 1, 7, 1, 0, 0, 0, loc) // Jan 7th 2022
	end := time.Date(2022, 1, 10, 1, 0, 0, 0, loc)  // Jan 10th 2022
	actual := RangeFromTimeStepper(start, end, 0, 0, 1, 0, 0, 0)
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

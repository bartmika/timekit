package timekit

import (
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

func TestIsFirstDayOfYear(t *testing.T) {
	loc := time.UTC                              // closure can be used if necessary
	t1 := time.Date(2022, 1, 7, 1, 0, 0, 0, loc) // Jan 7th
	t2 := time.Date(2022, 1, 1, 1, 0, 0, 0, loc) // Jan 1st

	isT1OK := IsFirstDayOfYear(t1)
	if isT1OK != false {
		t.Errorf("Incorrect boolean, got %v but was expecting %v", isT1OK, false)
	}
	isT2OK := IsFirstDayOfYear(t2)
	if isT2OK != true {
		t.Errorf("Incorrect boolean, got %v but was expecting %v", isT1OK, true)
	}
}

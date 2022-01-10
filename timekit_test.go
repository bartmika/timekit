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
		return time.Date(2000, 1, 1, 0, 0, 0, 0, loc)
	}

	actual := FirstDayOfLastYear(timeFn)
	expected := time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC)
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestFirstDayOfThisYear(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 7, 1, 0, 0, 0, 0, loc) // July 1st
	}

	actual := FirstDayOfThisYear(timeFn)
	expected := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestFirstDayOfLastMonth(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 1, 1, 0, 0, 0, 0, loc)
	}

	actual := FirstDayOfLastMonth(timeFn)
	expected := time.Date(1999, 12, 1, 0, 0, 0, 0, time.UTC)
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestFirstDayOfThisMonth(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 1, 20, 0, 0, 0, 0, loc) // Jan 20th
	}

	actual := FirstDayOfThisMonth(timeFn)
	expected := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestMidnight(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 1, 20, 5, 0, 0, 0, loc) // Jan 20th 5 AM
	}

	actual := Midnight(timeFn)
	expected := time.Date(2000, 1, 20, 0, 0, 0, 0, time.UTC)
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestMidnightTomorrow(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 1, 20, 5, 0, 0, 0, loc) // Jan 20th 5 AM
	}

	actual := MidnightTomorrow(timeFn)
	expected := time.Date(2000, 1, 21, 0, 0, 0, 0, time.UTC)
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestNoon(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 1, 20, 5, 0, 0, 0, loc) // Jan 20th 5 AM
	}

	actual := Noon(timeFn)
	expected := time.Date(2000, 1, 20, 12, 0, 0, 0, time.UTC)
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

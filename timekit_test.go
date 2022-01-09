package timekit

import (
	"testing"
	"time"
)

// Developers Note:
// Special thanks to the following article which helped explain how to stub
// out the `time.Now()` function in Golang.
// https://labs.yulrizka.com/en/stubbing-time-dot-now-in-golang/

func TestFirstDayFromLastYear(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 1, 1, 0, 0, 0, 0, loc)
	}

	actual := FirstDayFromLastYear(timeFn)
	expected := time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC)
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestFirstDayFromThisYear(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2000, 7, 1, 0, 0, 0, 0, loc) // July 7th
	}

	actual := FirstDayFromThisYear(timeFn)
	expected := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	if actual != expected {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

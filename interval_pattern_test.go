package timekit

import (
	"testing"
	"time"
)

func TestGetFutureDateByFiveMinuteIntervalPattern(t *testing.T) {
	loc := time.UTC                                    // closure can be used if necessary
	dt := time.Date(2022, 1, 9, 1, 0, 0, 0, loc)       // Sunday Jan 9th - 1:00 AM
	expected := time.Date(2022, 1, 9, 1, 0, 0, 0, loc) // Sunday Jan 9th - 1:00 AM
	actual := GetFutureDateByFiveMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 1, 0, 0, loc)       // Sunday Jan 9th - 1:01 AM
	expected = time.Date(2022, 1, 9, 1, 5, 0, 0, loc) // Sunday Jan 9th - 1:05 AM
	actual = GetFutureDateByFiveMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 28, 0, 0, loc)       // Sunday Jan 9th - 1:28 AM
	expected = time.Date(2022, 1, 9, 1, 30, 0, 0, loc) // Sunday Jan 9th - 1:30 AM
	actual = GetFutureDateByFiveMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 59, 0, 0, loc)      // Sunday Jan 9th - 1:59 AM
	expected = time.Date(2022, 1, 9, 2, 0, 0, 0, loc) // Sunday Jan 9th - 2:00 AM
	actual = GetFutureDateByFiveMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}
}

func TestGetFutureDateByTenMinuteIntervalPattern(t *testing.T) {
	loc := time.UTC                                    // closure can be used if necessary
	dt := time.Date(2022, 1, 9, 1, 0, 0, 0, loc)       // Sunday Jan 9th - 1:00 AM
	expected := time.Date(2022, 1, 9, 1, 0, 0, 0, loc) // Sunday Jan 9th - 1:00 AM
	actual := GetFutureDateByTenMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 1, 0, 0, loc)        // Sunday Jan 9th - 1:01 AM
	expected = time.Date(2022, 1, 9, 1, 10, 0, 0, loc) // Sunday Jan 9th - 1:05 AM
	actual = GetFutureDateByTenMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 28, 0, 0, loc)       // Sunday Jan 9th - 1:28 AM
	expected = time.Date(2022, 1, 9, 1, 30, 0, 0, loc) // Sunday Jan 9th - 1:30 AM
	actual = GetFutureDateByTenMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 59, 0, 0, loc)      // Sunday Jan 9th - 1:59 AM
	expected = time.Date(2022, 1, 9, 2, 0, 0, 0, loc) // Sunday Jan 9th - 2:00 AM
	actual = GetFutureDateByTenMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 23, 59, 0, 0, loc)      // Sunday Jan 9th - 23:59
	expected = time.Date(2022, 1, 10, 0, 0, 0, 0, loc) // Sunday Jan 10th - 00:00
	actual = GetFutureDateByTenMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}
}

func TestGetFutureDateByFiveteenMinuteIntervalPattern(t *testing.T) {
	loc := time.UTC                                    // closure can be used if necessary
	dt := time.Date(2022, 1, 9, 1, 0, 0, 0, loc)       // Sunday Jan 9th - 1:00 AM
	expected := time.Date(2022, 1, 9, 1, 0, 0, 0, loc) // Sunday Jan 9th - 1:00 AM
	actual := GetFutureDateByFiveteenMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 1, 0, 0, loc)        // Sunday Jan 9th - 1:01 AM
	expected = time.Date(2022, 1, 9, 1, 15, 0, 0, loc) // Sunday Jan 9th - 1:15 AM
	actual = GetFutureDateByFiveteenMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 28, 0, 0, loc)       // Sunday Jan 9th - 1:28 AM
	expected = time.Date(2022, 1, 9, 1, 30, 0, 0, loc) // Sunday Jan 9th - 1:30 AM
	actual = GetFutureDateByFiveteenMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 59, 0, 0, loc)      // Sunday Jan 9th - 1:59 AM
	expected = time.Date(2022, 1, 9, 2, 0, 0, 0, loc) // Sunday Jan 9th - 2:00 AM
	actual = GetFutureDateByFiveteenMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}
}

func TestGetFutureDateByThirtyMinuteIntervalPattern(t *testing.T) {
	loc := time.UTC                                    // closure can be used if necessary
	dt := time.Date(2022, 1, 9, 1, 0, 0, 0, loc)       // Sunday Jan 9th - 1:00 AM
	expected := time.Date(2022, 1, 9, 1, 0, 0, 0, loc) // Sunday Jan 9th - 1:00 AM
	actual := GetFutureDateByThirtyMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 1, 0, 0, loc)        // Sunday Jan 9th - 1:01 AM
	expected = time.Date(2022, 1, 9, 1, 30, 0, 0, loc) // Sunday Jan 9th - 1:30 AM
	actual = GetFutureDateByThirtyMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 28, 0, 0, loc)       // Sunday Jan 9th - 1:28 AM
	expected = time.Date(2022, 1, 9, 1, 30, 0, 0, loc) // Sunday Jan 9th - 1:30 AM
	actual = GetFutureDateByThirtyMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 59, 0, 0, loc)      // Sunday Jan 9th - 1:59 AM
	expected = time.Date(2022, 1, 9, 2, 0, 0, 0, loc) // Sunday Jan 9th - 2:00 AM
	actual = GetFutureDateByThirtyMinuteIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}
}

func TestGetFutureDateByOneHourIntervalPattern(t *testing.T) {
	loc := time.UTC                                    // closure can be used if necessary
	dt := time.Date(2022, 1, 9, 1, 0, 0, 0, loc)       // Sunday Jan 9th - 1:00 AM
	expected := time.Date(2022, 1, 9, 1, 0, 0, 0, loc) // Sunday Jan 9th - 1:00 AM
	actual := GetFutureDateByOneHourIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 1, 0, 0, loc)       // Sunday Jan 9th - 1:01 AM
	expected = time.Date(2022, 1, 9, 2, 0, 0, 0, loc) // Sunday Jan 9th - 2:00 AM
	actual = GetFutureDateByOneHourIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 28, 0, 0, loc)      // Sunday Jan 9th - 1:28 AM
	expected = time.Date(2022, 1, 9, 2, 0, 0, 0, loc) // Sunday Jan 9th - 2:00 AM
	actual = GetFutureDateByOneHourIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 1, 59, 0, 0, loc)      // Sunday Jan 9th - 1:59 AM
	expected = time.Date(2022, 1, 9, 2, 0, 0, 0, loc) // Sunday Jan 9th - 2:00 AM
	actual = GetFutureDateByOneHourIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	////////

	dt = time.Date(2022, 1, 9, 23, 59, 0, 0, loc)      // Sunday Jan 9th - 23:59
	expected = time.Date(2022, 1, 10, 0, 0, 0, 0, loc) // Sunday Jan 10th - 00:00
	actual = GetFutureDateByOneHourIntervalPattern(dt)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}
}

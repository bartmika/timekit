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

func TestYearsRange(t *testing.T) {
	loc := time.UTC // closure can be used if necessary
	t1 := time.Date(2000, 1, 1, 1, 0, 0, 0, loc)
	t2 := time.Date(2002, 1, 1, 1, 0, 0, 0, loc)

	actualYears := YearsRange(t1, t2)
	expectedYears := []int{2000, 2001, 2002}
	if reflect.DeepEqual(actualYears, expectedYears) == false {
		t.Errorf("Incorrect year ranges, got %v but was expecting %v", actualYears, expectedYears)
	}
}

func TestMonthsRange(t *testing.T) {
	loc := time.UTC // closure can be used if necessary

	// CASE 1 - Beginning of year.

	t1 := time.Date(2000, 1, 1, 1, 0, 0, 0, loc)
	t2 := time.Date(2000, 4, 1, 1, 0, 0, 0, loc)

	actualMonths := MonthRange(t1, t2)
	expectedMonths := []int{1, 2, 3, 4}
	if reflect.DeepEqual(actualMonths, expectedMonths) == false {
		t.Errorf("Incorrect month ranges, got %v but was expecting %v", actualMonths, expectedMonths)
	}

	// CASE 2 - In middle of year.

	t3 := time.Date(2000, 6, 1, 1, 0, 0, 0, loc)
	t4 := time.Date(2000, 12, 1, 1, 0, 0, 0, loc)

	actualMonths = MonthRange(t3, t4)
	expectedMonths = []int{6, 7, 8, 9, 10, 11, 12}
	if reflect.DeepEqual(actualMonths, expectedMonths) == false {
		t.Errorf("Incorrect month ranges, got %v but was expecting %v", actualMonths, expectedMonths)
	}

	// Case 3 - Between years.

	t5 := time.Date(2000, 11, 1, 1, 0, 0, 0, loc)
	t6 := time.Date(2001, 2, 1, 1, 0, 0, 0, loc)

	actualMonths = MonthRange(t5, t6)
	expectedMonths = []int{11, 12, 1, 2}
	if reflect.DeepEqual(actualMonths, expectedMonths) == false {
		t.Errorf("Incorrect month ranges, got %v but was expecting %v", actualMonths, expectedMonths)
	}
}

func TestDaysRange(t *testing.T) {
	loc := time.UTC // closure can be used if necessary

	// CASE 1 - Beginning of month.

	t1 := time.Date(2000, 1, 1, 1, 0, 0, 0, loc)
	t2 := time.Date(2000, 1, 10, 1, 0, 0, 0, loc)

	actualDays := DaysRange(t1, t2)
	expectedDays := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if reflect.DeepEqual(actualDays, expectedDays) == false {
		t.Errorf("Incorrect year ranges, got %v but was expecting %v", actualDays, expectedDays)
	}

	// Case 2 - Middle of month.

	t3 := time.Date(2000, 1, 15, 1, 0, 0, 0, loc)
	t4 := time.Date(2000, 1, 25, 1, 0, 0, 0, loc)

	actualDays = DaysRange(t3, t4)
	expectedDays = []int{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	if reflect.DeepEqual(actualDays, expectedDays) == false {
		t.Errorf("Incorrect year ranges, got %v but was expecting %v", actualDays, expectedDays)
	}

	// Case 3 - Across months.

	t5 := time.Date(2000, 1, 29, 1, 0, 0, 0, loc)
	t6 := time.Date(2000, 2, 2, 1, 0, 0, 0, loc)

	actualDays = DaysRange(t5, t6)
	expectedDays = []int{29, 30, 31, 1, 2}
	if reflect.DeepEqual(actualDays, expectedDays) == false {
		t.Errorf("Incorrect year ranges, got %v but was expecting %v", actualDays, expectedDays)
	}
}

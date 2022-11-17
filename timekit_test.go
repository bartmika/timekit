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

func TestGetWeekNumberFromDate(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary

	// Week 52.

	t1 := time.Date(2022, 1, 1, 1, 0, 0, 0, loc) // Friday Jan 1st.
	wkn := GetWeekNumberFromDate(t1)
	if wkn != 52 {
		t.Errorf("Incorrect integer, got %v but was expecting %v", wkn, 52)
	}

	// Week 1.

	t2 := time.Date(2022, 1, 7, 1, 0, 0, 0, loc) // Friday Jan 7th
	wkn = GetWeekNumberFromDate(t2)
	if wkn != 1 {
		t.Errorf("Incorrect integer, got %v but was expecting %v", wkn, 1)
	}

	// Week 52 (again).

	t3 := time.Date(2022, 12, 30, 1, 0, 0, 0, loc) // December 30th
	wkn = GetWeekNumberFromDate(t3)
	if wkn != 52 {
		t.Errorf("Incorrect integer, got %v but was expecting %v", wkn, 52)
	}
}

func TestGetFirstDateFromWeekAndYear(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary

	// Week 1

	expected := time.Date(2022, 1, 3, 1, 0, 0, 0, loc)
	actual := GetFirstDateFromWeekAndYear(1, 2022, loc)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	// Week 4

	expected = time.Date(2022, 1, 24, 1, 0, 0, 0, loc)
	actual = GetFirstDateFromWeekAndYear(4, 2022, loc)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	// Week 52 (from last week)

	expected = time.Date(2022, 1, 1, 1, 0, 0, 0, loc)
	actual = GetFirstDateFromWeekAndYear(52, 2022, loc)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}
}

func TestGetFirstDateFromMonthAndYear(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary

	// Week 1

	expected := time.Date(2022, 1, 1, 1, 0, 0, 0, loc)
	actual := GetFirstDateFromMonthAndYear(1, 2022, loc)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}
}

func TestGetWeekNumberFromTotalDaysCount(t *testing.T) {
	type ExpectedTestResult struct {
		StartDay uint64
		EndDay   uint64
		Week     uint64
	}
	testData := []*ExpectedTestResult{
		{0, 0, 0}, // This is a zero case to test for.

		// DEVELOPERS NOTE: Here are some hard-coded test values. Uncomment these
		// if you'd like to run for human-entered values instead of the for loop
		// below.
		// {1, 7, 1}, {8, 14, 2}, {15, 21, 3}, {22, 28, 4}, {29, 35, 5},
		// {36, 42, 6}, {43, 49, 7},
	}

	// The following code will populate our test weeks to verify our code works.
	days := uint64(1)
	for wk := uint64(1); wk <= 52; wk++ {
		testDatum := &ExpectedTestResult{days, days + 6, wk}
		days += 7
		testData = append(testData, testDatum)
	}

	// Iterate through all the test data on our function and verify that the
	// function works correctly.
	for _, testDatum := range testData {
		for d := uint64(testDatum.StartDay); d <= testDatum.EndDay; d++ {
			wk := GetWeekNumberFromTotalDaysCount(d)
			if wk != testDatum.Week {
				t.Errorf("Incorrect week number for day #%d, got week %d but was expecting week %d", d, wk, testDatum.Week)
				return
			}
		}
	}
}

func TestGetDayOfWeekUsingTomohikoSakamotoAlgorithm(t *testing.T) {
	// Value of `1` for input of `14/09/1998` was taken from this link:
	// for https://www.geeksforgeeks.org/find-day-of-the-week-for-a-given-date/
	actualDay := GetDayOfWeekUsingTomohikoSakamotoAlgorithm(14, 9, 1998)
	if actualDay != 1 {
		t.Errorf("Incorrect actual day #%d, expecting day %d", actualDay, 1)
	}

	// Returned value of `2` for input `15/11/2022` was taken by looking in the
	// calendar.
	actualDay = GetDayOfWeekUsingTomohikoSakamotoAlgorithm(15, 11, 2022)
	if actualDay != 2 {
		t.Errorf("Incorrect actual day #%d, expecting day %d", actualDay, 1)
	}
	// Used calendar to for test.
	actualDay = GetDayOfWeekUsingTomohikoSakamotoAlgorithm(1, 1, 2023)
	if actualDay != 0 { // 0 = Sunday
		t.Errorf("Incorrect actual day #%d, expecting day %d", actualDay, 1)
	}
}

func TestAddWeeksToTime(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary

	// Week 1

	expected := time.Date(2022, 1, 8, 1, 0, 0, 0, loc)
	dt := time.Date(2022, 1, 1, 1, 0, 0, 0, loc)
	actual := AddWeeksToTime(dt, 1)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	// Week 2

	expected = time.Date(2022, 1, 15, 1, 0, 0, 0, loc)
	dt = time.Date(2022, 1, 1, 1, 0, 0, 0, loc)
	actual = AddWeeksToTime(dt, 2)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}

	// Week 3

	expected = time.Date(2022, 1, 22, 1, 0, 0, 0, loc)
	dt = time.Date(2022, 1, 1, 1, 0, 0, 0, loc)
	actual = AddWeeksToTime(dt, 3)
	if expected != actual {
		t.Errorf("Incorrect date, got %v but was expecting %v", actual, expected)
	}
}

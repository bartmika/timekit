package timekit

import (
	"reflect"
	"testing"
	"time"
)

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

func TestWeeksRange(t *testing.T) {
	loc := time.UTC // closure can be used if necessary

	t1 := time.Date(2022, 1, 1, 1, 0, 0, 0, loc)  // Week 52
	t2 := time.Date(2022, 1, 10, 1, 0, 0, 0, loc) // Week 2

	actualDays := WeeksRange(t1, t2)
	expectedDays := []int{52, 1, 2} // Weeks 52, 1, 2
	if reflect.DeepEqual(actualDays, expectedDays) == false {
		t.Errorf("Incorrect year ranges, got %v but was expecting %v", actualDays, expectedDays)
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

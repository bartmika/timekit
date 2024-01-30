package timekit

import (
	"testing"
	"time"
)

func TestHourlyRangeForTime(t *testing.T) {
	loc := time.UTC                                    // closure can be used if necessary
	given := time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	dtr := HourlyRangeForTime(given)
	exp1 := time.Date(2023, 12, 18, 9, 0, 0, 0, loc)  // Monday Dec 18th - 9 AM
	exp2 := time.Date(2023, 12, 18, 10, 0, 0, 0, loc) // Monday Dec 18th - 10 AM
	if exp1 != dtr.Start {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
	}
	if exp2 != dtr.End {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
	}
}

func TestHourlyRangeForNow(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	}

	dtr := HourlyRangeForNow(timeFn)
	exp1 := time.Date(2023, 12, 18, 9, 0, 0, 0, loc)  // Monday Dec 18th - 9 AM
	exp2 := time.Date(2023, 12, 18, 10, 0, 0, 0, loc) // Monday Dec 18th - 10 AM
	if exp1 != dtr.Start {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
	}
	if exp2 != dtr.End {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
	}
}

func TestHourlyRangesBetweenTimes(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC                                    // closure can be used if necessary
	start := time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	end := time.Date(2023, 12, 18, 11, 30, 0, 0, loc)  // Monday Dec 18th - 11:30 AM

	dtrdtr := HourlyRangesBetweenTimes(start, end)
	for i, dtr := range dtrdtr {
		switch i {
		case 0:
			exp1 := time.Date(2023, 12, 18, 9, 0, 0, 0, loc)  // Monday Dec 18th - 9 AM
			exp2 := time.Date(2023, 12, 18, 10, 0, 0, 0, loc) // Monday Dec 18th - 10 AM
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		case 1:
			exp1 := time.Date(2023, 12, 18, 10, 0, 0, 0, loc) // Monday Dec 18th - 10 AM
			exp2 := time.Date(2023, 12, 18, 11, 0, 0, 0, loc) // Monday Dec 18th - 11 AM
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		case 2:
			exp1 := time.Date(2023, 12, 18, 11, 0, 0, 0, loc) // Monday Dec 18th - 11 AM
			exp2 := time.Date(2023, 12, 18, 12, 0, 0, 0, loc) // Monday Dec 18th - 12 AM
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		}
	}
}

func TestDailyRangeForTime(t *testing.T) {
	loc := time.UTC                                    // closure can be used if necessary
	given := time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	dtr := DailyRangeForTime(given)
	exp1 := time.Date(2023, 12, 18, 0, 0, 0, 0, loc) // Monday Dec 18th - 12 AM
	exp2 := time.Date(2023, 12, 19, 0, 0, 0, 0, loc) // Monday Dec 19th - 12 AM
	if exp1 != dtr.Start {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
	}
	if exp2 != dtr.End {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
	}
}

func TestDailyRangeForNow(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	}

	dtr := DailyRangeForNow(timeFn)
	exp1 := time.Date(2023, 12, 18, 0, 0, 0, 0, loc) // Monday Dec 18th - 12 AM
	exp2 := time.Date(2023, 12, 19, 0, 0, 0, 0, loc) // Monday Dec 19th - 12 AM
	if exp1 != dtr.Start {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
	}
	if exp2 != dtr.End {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
	}
}

func TestDailyRangesBetweenTimes(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC                                    // closure can be used if necessary
	start := time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	end := time.Date(2023, 12, 20, 11, 30, 0, 0, loc)  // Monday Dec 20th - 11:30 AM

	dtrdtr := DailyRangesBetweenTimes(start, end)
	for i, dtr := range dtrdtr {
		switch i {
		case 0:
			exp1 := time.Date(2023, 12, 18, 0, 0, 0, 0, loc) // Monday Dec 18th - 9 AM
			exp2 := time.Date(2023, 12, 19, 0, 0, 0, 0, loc) // Monday Dec 18th - 10 AM
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		case 1:
			exp1 := time.Date(2023, 12, 19, 0, 0, 0, 0, loc) // Monday Dec 18th - 10 AM
			exp2 := time.Date(2023, 12, 20, 0, 0, 0, 0, loc) // Monday Dec 18th - 11 AM
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		case 2:
			exp1 := time.Date(2023, 12, 20, 0, 0, 0, 0, loc) // Monday Dec 18th - 11 AM
			exp2 := time.Date(2023, 12, 21, 0, 0, 0, 0, loc) // Monday Dec 18th - 12 AM
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		}
	}
}

func TestISOWeeklyRangeForTime(t *testing.T) {
	loc := time.UTC                                    // closure can be used if necessary
	given := time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	dtr := ISOWeeklyRangeForTime(given)
	exp1 := time.Date(2023, 12, 18, 0, 0, 0, 0, loc) // Monday Dec 18th - 12 AM
	exp2 := time.Date(2023, 12, 24, 0, 0, 0, 0, loc) // Monday Dec 24th - 12 AM
	if exp1 != dtr.Start {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
	}
	if exp2 != dtr.End {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
	}
}

func TestISOWeeklyRangeForNow(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	}

	dtr := ISOWeeklyRangeForNow(timeFn)
	exp1 := time.Date(2023, 12, 18, 0, 0, 0, 0, loc) // Monday Dec 18th - 12 AM
	exp2 := time.Date(2023, 12, 24, 0, 0, 0, 0, loc) // Monday Dec 24th - 12 AM
	if exp1 != dtr.Start {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
	}
	if exp2 != dtr.End {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
	}
}

func TestISOWeeklyRangesBetweenTimes(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC                                    // closure can be used if necessary
	start := time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	end := time.Date(2024, 01, 01, 11, 30, 0, 0, loc)  // Monday Dec 20th - 11:30 AM

	dtrdtr := ISOWeeklyRangesBetweenTimes(start, end)
	for i, dtr := range dtrdtr {
		switch i {
		case 0:
			exp1 := time.Date(2023, 12, 18, 0, 0, 0, 0, loc) // Monday Dec 18th - 12 AM
			exp2 := time.Date(2023, 12, 24, 0, 0, 0, 0, loc) // Monday Dec 24th - 12 AM
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		case 1:
			exp1 := time.Date(2023, 12, 25, 0, 0, 0, 0, loc)
			exp2 := time.Date(2023, 12, 31, 0, 0, 0, 0, loc)
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		case 2:
			exp1 := time.Date(2024, 01, 01, 0, 0, 0, 0, loc)
			exp2 := time.Date(2024, 01, 07, 0, 0, 0, 0, loc)
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		}
	}
}

func TestMonthlyRangeForTime(t *testing.T) {
	loc := time.UTC                                    // closure can be used if necessary
	given := time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	dtr := MonthlyRangeForTime(given)
	exp1 := time.Date(2023, 12, 01, 0, 0, 0, 0, loc)
	exp2 := time.Date(2024, 01, 01, 0, 0, 0, 0, loc)
	if exp1 != dtr.Start {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
	}
	if exp2 != dtr.End {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
	}
}

func TestMonthlyRangeForNow(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	}

	dtr := MonthlyRangeForNow(timeFn)
	exp1 := time.Date(2023, 12, 01, 0, 0, 0, 0, loc)
	exp2 := time.Date(2024, 01, 01, 0, 0, 0, 0, loc)
	if exp1 != dtr.Start {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
	}
	if exp2 != dtr.End {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
	}
}

func TestMonthlyRangesBetweenTimes(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC                                    // closure can be used if necessary
	start := time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	end := time.Date(2024, 02, 01, 01, 30, 0, 0, loc)

	dtrdtr := MonthlyRangesBetweenTimes(start, end)
	for i, dtr := range dtrdtr {
		switch i {
		case 0:
			exp1 := time.Date(2023, 12, 01, 0, 0, 0, 0, loc)
			exp2 := time.Date(2024, 01, 01, 0, 0, 0, 0, loc)
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		case 1:
			exp1 := time.Date(2024, 01, 01, 0, 0, 0, 0, loc)
			exp2 := time.Date(2024, 02, 01, 0, 0, 0, 0, loc)
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		case 2:
			exp1 := time.Date(2024, 12, 18, 11, 0, 0, 0, loc) // Monday Dec 18th - 11 AM
			exp2 := time.Date(2024, 12, 18, 12, 0, 0, 0, loc) // Monday Dec 18th - 12 AM
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		}
	}
}

func TestYearlyRangeForTime(t *testing.T) {
	loc := time.UTC                                    // closure can be used if necessary
	given := time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	dtr := YearlyRangeForTime(given)
	exp1 := time.Date(2023, 01, 01, 0, 0, 0, 0, loc)
	exp2 := time.Date(2024, 01, 01, 0, 0, 0, 0, loc)
	if exp1 != dtr.Start {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
	}
	if exp2 != dtr.End {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
	}
}

func TestYearlyRangeForNow(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC // closure can be used if necessary
	timeFn := func() time.Time {
		return time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	}

	dtr := YearlyRangeForNow(timeFn)
	exp1 := time.Date(2023, 01, 01, 0, 0, 0, 0, loc)
	exp2 := time.Date(2024, 01, 01, 0, 0, 0, 0, loc)
	if exp1 != dtr.Start {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
	}
	if exp2 != dtr.End {
		t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
	}
}

func TestYearlyRangesBetweenTimes(t *testing.T) {
	// Stub out the `time.Now()` function with our custom value so we can
	// simulate being in this current data.
	loc := time.UTC                                    // closure can be used if necessary
	start := time.Date(2023, 12, 18, 9, 30, 0, 0, loc) // Monday Dec 18th - 9:30 AM
	end := time.Date(2025, 12, 18, 11, 30, 0, 0, loc)

	dtrdtr := YearlyRangesBetweenTimes(start, end)
	for i, dtr := range dtrdtr {
		switch i {
		case 0:
			exp1 := time.Date(2023, 01, 01, 0, 0, 0, 0, loc)
			exp2 := time.Date(2024, 01, 01, 0, 0, 0, 0, loc)
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		case 1:
			exp1 := time.Date(2024, 01, 01, 0, 0, 0, 0, loc)
			exp2 := time.Date(2025, 01, 01, 0, 0, 0, 0, loc)
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		case 2:
			exp1 := time.Date(2025, 01, 01, 0, 0, 0, 0, loc)
			exp2 := time.Date(2026, 01, 01, 0, 0, 0, 0, loc)
			if exp1 != dtr.Start {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.Start, exp1)
			}
			if exp2 != dtr.End {
				t.Errorf("Incorrect date, got %s but was expecting %s", dtr.End, exp2)
			}
			break
		}
	}
}

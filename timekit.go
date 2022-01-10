package timekit

import "time"

// FirstDateFromLastYear returns first date (with 0:00 hour) from last year.
func FirstDateFromLastYear(now func() time.Time) time.Time {
	dt := now().UTC()
	lastYear := dt.Year() - 1
	return time.Date(lastYear, 1, 1, 0, 0, 0, 0, time.UTC)
}

// FirstDateFromThisYear returns the date (with 0:00 hour) from the first date of this year.
func FirstDateFromThisYear(now func() time.Time) time.Time {
	dt := now().UTC()
	lastYear := dt.Year()
	return time.Date(lastYear, 1, 1, 0, 0, 0, 0, time.UTC)
}

// FirstDateFromLastMonth returns the date (with 0:00 hour) of the first day from last month.
func FirstDateFromLastMonth(now func() time.Time) time.Time {
	dt := now().UTC()
	return time.Date(dt.Year(), dt.Month()-1, 1, 0, 0, 0, 0, time.UTC)
}

// FirstDateFromThisMonth returns the first date (with 0:00 hour) from this month.
func FirstDateFromThisMonth(now func() time.Time) time.Time {
	dt := now().UTC()
	return time.Date(dt.Year(), dt.Month(), 1, 0, 0, 0, 0, time.UTC)
}

// Midnight will return today's date at 12 o’clock (or 0:00) during the night.
func Midnight(now func() time.Time) time.Time {
	dt := now().UTC()
	return time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, time.UTC)
}

// Noon will return today's date at 12 o'clock (or 12:00) during the day.
func Noon(now func() time.Time) time.Time {
	dt := now().UTC()
	return time.Date(dt.Year(), dt.Month(), dt.Day(), 12, 0, 0, 0, time.UTC)
}

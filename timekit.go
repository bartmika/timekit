package timekit

import "time"

// FirstDayOfLastYear returns first date (with 0:00 hour) from last year.
func FirstDayOfLastYear(now func() time.Time) time.Time {
	dt := now()
	lastYear := dt.Year() - 1
	return time.Date(lastYear, 1, 1, 0, 0, 0, 0, dt.Location())
}

// FirstDayOfThisYear returns the date (with 0:00 hour) from the first date of this year.
func FirstDayOfThisYear(now func() time.Time) time.Time {
	dt := now()
	lastYear := dt.Year()
	return time.Date(lastYear, 1, 1, 0, 0, 0, 0, dt.Location())
}

// FirstDayOfLastMonth returns the date (with 0:00 hour) of the first day from last month.
func FirstDayOfLastMonth(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month()-1, 1, 0, 0, 0, 0, dt.Location())
}

// FirstDayOfThisMonth returns the first date (with 0:00 hour) from this month.
func FirstDayOfThisMonth(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month(), 1, 0, 0, 0, 0, dt.Location())
}

// Midnight will return today's date at 12 o’clock (or 0:00) during the night.
func Midnight(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
}

// Midnight will return tomorrows date at 12 o’clock (or 0:00) during the night.
func MidnightTomorrow(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month(), dt.Day()+1, 0, 0, 0, 0, dt.Location())
}

// Noon will return today's date at 12 o'clock (or 12:00) during the day.
func Noon(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month(), dt.Day(), 12, 0, 0, 0, dt.Location())
}

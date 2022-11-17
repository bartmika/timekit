package timekit

import (
	"time"
)

// FirstDayOfLastYear returns first date (with 0:00 hour) from last calendar year.
func FirstDayOfLastYear(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year()-1, 1, 1, 0, 0, 0, 0, dt.Location())
}

// FirstDayOfThisYear returns the date (with 0:00 hour) from the first date of this calendar year.
func FirstDayOfThisYear(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), 1, 1, 0, 0, 0, 0, dt.Location())
}

// FirstDayOfNextYear returns date (12AM) of the first date of next calendar year.
func FirstDayOfNextYear(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year()+1, 1, 1, 0, 0, 0, 0, dt.Location())
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

// FirstDayOfNextMonth returns next months first day (in 12 AM hours).
func FirstDayOfNextMonth(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month()+1, 1, 0, 0, 0, 0, dt.Location())
}

// MidnightYesterday return 12 AM date of yesterday.
func MidnightYesterday(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month(), dt.Day()-1, 0, 0, 0, 0, dt.Location())
}

// Midnight return today's date at 12 o’clock (or 0:00) during the night.
func Midnight(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
}

// MidnightTomorrow will return tomorrows date at 12 o’clock (or 0:00) during the night.
func MidnightTomorrow(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month(), dt.Day()+1, 0, 0, 0, 0, dt.Location())
}

// Noon will return today's date at 12 o'clock (or 12:00) during the day.
func Noon(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month(), dt.Day(), 12, 0, 0, 0, dt.Location())
}

// FirstDayOfLastISOWeek returns the previous week's monday date.
func FirstDayOfLastISOWeek(now func() time.Time) time.Time {
	dt := now()

	// iterate back to Monday
	for dt.Weekday() != time.Monday {
		dt = dt.AddDate(0, 0, -1)
	}
	dt = dt.AddDate(0, 0, -1) // Skip the current monday we are on!

	// iterate ONCE AGAIN back to Monday
	for dt.Weekday() != time.Monday {
		dt = dt.AddDate(0, 0, -1)
	}

	return dt
}

// FirstDayOfThisISOWeek return monday's date of this week. Please note monday is considered the first day of the week according to ISO 8601 and not sunday (which is what is used in Canada and USA).
func FirstDayOfThisISOWeek(now func() time.Time) time.Time {
	dt := now()

	// iterate back to Monday
	for dt.Weekday() != time.Monday {
		dt = dt.AddDate(0, 0, -1)
	}

	return dt
}

// LastDayOfThisISOWeek return sunday's date of this week. Please note sunday is considered the last day of the week according to ISO 8601.
func LastDayOfThisISOWeek(now func() time.Time) time.Time {
	dt := now()

	// iterate forward to Sunday
	for dt.Weekday() != time.Sunday {
		dt = dt.AddDate(0, 0, 1)
	}

	return dt
}

// FirstDayOfNextISOWeek return date of the upcoming monday.
func FirstDayOfNextISOWeek(now func() time.Time) time.Time {
	dt := now()

	// iterate forward to next Monday
	for dt.Weekday() != time.Monday {
		dt = dt.AddDate(0, 0, 1)
	}

	return dt
}

// IsFirstDayOfYear returns true or false depending on whether the date inputted falls on the very first day of the year.
func IsFirstDayOfYear(dt time.Time) bool {
	return dt.Day() == 1 && dt.Month() == 1
}

// GetWeekNumberFromDate will return the week number for the inputted date.
func GetWeekNumberFromDate(dt time.Time) int {
	_, week := dt.ISOWeek()
	return week
}

// GetFirstDateFromWeekAndYear returns the first date for the particular week in the year inputted.
func GetFirstDateFromWeekAndYear(wk int, year int, loc *time.Location) time.Time {
	start := time.Date(year, 1, 1, 1, 0, 0, 0, loc)
	end := time.Date(year+1, 1, 1, 1, 0, 0, 0, loc)

	ts := NewTimeStepper(start, end, 0, 0, 1, 0, 0, 0) // Step by day.
	dt := ts.Get()                                     // Get first day.

	// Please note, there may be cases where week 52 happens in January,
	// the the docs via https://pkg.go.dev/time#Time.ISOWeek.

	// CASE 1
	week := GetWeekNumberFromDate(dt)
	if week == wk {
		return dt
	}

	// CASE 2
	for ts.Next() {
		dt = ts.Get()
		week := GetWeekNumberFromDate(dt)
		if week == wk {
			break
		}
	}
	return dt
}

// GetFirstDateFromMonthAndYear returns the first day in the month/year specified.
func GetFirstDateFromMonthAndYear(month int, year int, loc *time.Location) time.Time {
	return time.Date(year, time.Month(month), 1, 1, 0, 0, 0, loc)
}

// GetWeekNumberFromTotalDaysCount returns the week number from total days count. For example daysCount=8, returns=2 or daysCount=365, returned=52.
func GetWeekNumberFromTotalDaysCount(daysCount uint64) uint64 {
	if daysCount == 0 {
		return 0
	}

	wk := uint64(1)
	wkDay := uint64(0)
	for day := uint64(0); day < daysCount; day++ {
		if wkDay == 7 { // End of week detected
			wkDay = 1 // Reset day of week counter to the beginning of the week.
			wk++      // Store the fact that we are starting a new week.
		} else {
			wkDay++
		}
	}
	return wk
}

// GetDayOfWeekUsingTomohikoSakamotoAlgorithm returns the day of the week where `0` = Sunday, `1` = Monday, etc.
func GetDayOfWeekUsingTomohikoSakamotoAlgorithm(d uint64, m uint64, y uint64) uint64 {
	// Please see `Sakamoto's methods` via https://en.wikipedia.org/wiki/Determination_of_the_day_of_the_week.
	a := []uint64{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}
	if m < 3 {
		y--
	}
	day := (y + y/4 - y/100 + y/400 + a[m-1] + d) % 7
	return day
}

// AddWeeksToTime returns new time with weeks added to it.
func AddWeeksToTime(dt time.Time, weeks int) time.Time {
	return dt.AddDate(0, 0, 7*weeks)
}

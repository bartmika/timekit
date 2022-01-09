package timekit

import "time"

// FirstDayFromLastYear will return the `time.Time` struct of the first day
// (zero hour) from last year in UTC timezone.
func FirstDayFromLastYear(now func() time.Time) time.Time {
	dt := now().UTC()
	lastYear := dt.Year() - 1
	return time.Date(lastYear, 1, 1, 0, 0, 0, 0, time.UTC)
}

// FirstDayFromThisYear will return the `time.Time` struct of the first day
// (zero hour) from this year in UTC timezone.
func FirstDayFromThisYear(now func() time.Time) time.Time {
	dt := now().UTC()
	lastYear := dt.Year()
	return time.Date(lastYear, 1, 1, 0, 0, 0, 0, time.UTC)
}

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

	return time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location()) // Get midnight of this day.
}

// FirstDayOfThisISOWeek return monday's date of this week. Please note monday is considered the first day of the week according to ISO 8601 and not sunday (which is what is used in Canada and USA).
func FirstDayOfThisISOWeek(now func() time.Time) time.Time {
	dt := now()

	// iterate back to Monday
	for dt.Weekday() != time.Monday {
		dt = dt.AddDate(0, 0, -1)
	}

	return time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location()) // Get midnight of this day.
}

// LastDayOfThisISOWeek return sunday's date of this week. Please note sunday is considered the last day of the week according to ISO 8601.
func LastDayOfThisISOWeek(now func() time.Time) time.Time {
	dt := now()

	// iterate forward to Sunday
	for dt.Weekday() != time.Sunday {
		dt = dt.AddDate(0, 0, 1)
	}

	return time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location()) // Get midnight of this day.
}

// FirstDayOfNextISOWeek return date of the upcoming monday.
func FirstDayOfNextISOWeek(now func() time.Time) time.Time {
	startAt := now()

	// CASE 1 OF 2: Today is monday, therefore simply add seven to current
	//              date/time and return.
	if startAt.Weekday() == time.Monday {
		// Generate new date 7 days from now to start at the next monday.
		return time.Date(startAt.Year(), startAt.Month(), startAt.Day()+7, 0, 0, 0, 0, startAt.Location())
	}
	// Case 2 of 2: Today is not monday so we must find the next instance of it.

	// Variable holds additional 7 days to current start date so we have an
	// extra week to look through.
	endAt := now().AddDate(0, 0, 7)

	// Create a time-stepper go iterate per day from today to finish date and
	// then stop at monday.
	ts := NewTimeStepper(startAt, endAt, 0, 0, 1, 0, 0, 0)

	var dt time.Time

	// Itere through all the steps until you land on a `monday`.
	for ts.Next() {
		dt = ts.Get()
		if dt.Weekday() == time.Monday {
			break
		}
	}

	// Return the midnight time of the current date.
	return time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
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

// GetDatesForWeekdaysBetweenRange returns all the date-times between two dates that fall for the specific picked weekdays.
func GetDatesForWeekdaysBetweenRange(start time.Time, end time.Time, weekdays []int8) []time.Time {
	times := []time.Time{}

	// Create all the dates, incremented by day, from the start to finish datetimes.
	dts := RangeFromTimeStepper(start, end, 0, 0, 1, 0, 0, 0)

	// Iterate through all the datetimes and check to see if the weekdays match.
	for _, dt := range dts {
		wkInt := int8(dt.Weekday())

		// Iterate through the specificed weekdays.
		for _, weekday := range weekdays {
			if wkInt == weekday {
				times = append(times, dt)
			}
		}
	}
	return times
}

func IsTimeOnFirstWeekOfMonth(pickedDT time.Time) bool {
	currentYear, currentMonth, _ := pickedDT.Date()
	currentLocation := pickedDT.Location()

	firstDayOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)

	// Create all the dates, incremented by day, from the start to finish datetimes.
	dts := RangeFromTimeStepper(firstDayOfMonth, lastDayOfMonth, 0, 0, 1, 0, 0, 0)

	var isOnFirstWeek bool = false
	for _, dt := range dts {
		// Get weekday integer from the current iteration date.
		weekdayInt := int8(dt.Weekday())

		// For debugging purposes only.
		// log.Println(pickedDT, "|", dt, "|", weekdayInt, "|", pickedDT.Day(), "|", dt.Day())

		// Check if the picked day falls into the range of our first week.
		if pickedDT.Day() == dt.Day() {
			isOnFirstWeek = true
			break
		}

		if weekdayInt == 6 { // Last ISO week day so give up and stop loop.
			break
		}
	}

	return isOnFirstWeek
}

func IsTimeOnLastWeekOfMonth(pickedDT time.Time) bool {
	currentYear, currentMonth, _ := pickedDT.Date()
	currentLocation := pickedDT.Location()

	firstDayOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)
	lastDay := lastDayOfMonth.Day()

	var isOnLastWeek bool = false

	// Start backwards and try to find the first Sunday to terminate or matching.
	for day := lastDay; day >= 1; day-- {
		dt := time.Date(currentYear, currentMonth, day, 0, 0, 0, 0, currentLocation)

		// Get weekday integer from the current iteration date.
		weekdayInt := int8(dt.Weekday())

		if dt.Day() == pickedDT.Day() {
			isOnLastWeek = true
			break
		}

		// If we are entering an older week then stop the loop, we did not find anything.
		if weekdayInt == 0 {
			break
		}
	}

	return isOnLastWeek
}

// GetDatesByWeeklyBasedRecurringSchedule Generates a list of datetimes based on a weekly recuring schedule. Please note that dates start in first week and then week frequency is applied to restrict some weeks.
func GetDatesByWeeklyBasedRecurringSchedule(startDT time.Time, weekdays []int8, totalWeeks int, weekFrequency int) []time.Time {
	// Variable will store the datetimes that belong in the weekly based recurring schedule.
	scheduleTimes := []time.Time{}

	// Variable will calculate the last date based on total weeks in schedule.
	endDT := AddWeeksToTime(startDT, totalWeeks-1)

	// Create all the dates, incremented by day, from the start to finish datetimes.
	dts := RangeFromTimeStepper(startDT, endDT, 0, 0, 1, 0, 0, 0)

	// Variable is used to track when to trigger creation of the scheduled time.
	var weekFrequencyIterator int = 0

	// Variable used to track how many days have elapsed.
	var dayIterator int = 0

	// Iterate through all the days between the start to end date.
	for _, todayDT := range dts {
		// Get weekday integer from the current iteration date.
		weekdayInt := int8(todayDT.Weekday())

		// If seven days elapsed then the end of the week occured.
		if dayIterator%7 == 0 && dayIterator != 0 {
			weekFrequencyIterator++
		}

		// Iterate through the picked weekdays that the user has chosen and see
		// if today's date
		for _, weekday := range weekdays {
			if weekdayInt == weekday {
				if weekFrequencyIterator%weekFrequency == 0 {
					scheduleTimes = append(scheduleTimes, todayDT)
				}
			}
		}

		// We finished processing this day so keep track in our day iterator and continue.
		dayIterator++
	}
	return scheduleTimes
}

// GetDatesForExactDayByMonthlyBasedRecurringSchedule Generates a list of datetimes based on a monthly recuring schedule for the specific day number.
func GetDatesForExactDayByMonthlyBasedRecurringSchedule(startDT time.Time, totalMonths int, onExactDay int) []time.Time {
	// Variable will store the datetimes that belong in the monthly based recurring schedule.
	scheduleTimes := []time.Time{}

	// Variable will calculate the last date based on total weeks in schedule.
	endDT := startDT.AddDate(0, totalMonths-1, 0)

	// Create all the dates, incremented by day, from the start to finish datetimes.
	dts := RangeFromTimeStepper(startDT, endDT, 0, 0, 1, 0, 0, 0)

	// Iterate through all the days between the start to end date.
	for _, todayDT := range dts {
		if onExactDay == todayDT.Day() {
			scheduleTimes = append(scheduleTimes, todayDT)
		}
	}

	return scheduleTimes
}

// GetDatesForFirstWeekDayByMonthlyBasedRecurringSchedule Generates a list of datetimes based on a monthly recuring schedule which will find all the dates that fall on the weekday of the first week.
func GetDatesForFirstWeekDayByMonthlyBasedRecurringSchedule(startDT time.Time, totalMonths int, onFirstWeekday int) []time.Time {
	// Variable will store the datetimes that belong in the monthly based recurring schedule.
	scheduleTimes := []time.Time{}

	// Variable will calculate the last date based on total weeks in schedule.
	endDT := startDT.AddDate(0, totalMonths-1, 7) // Why "7"? Just to take into account extra week so we can handle for that particular months first week.

	// Create all the dates, incremented by day, from the start to finish datetimes.
	dts := RangeFromTimeStepper(startDT, endDT, 0, 0, 1, 0, 0, 0)

	monthIterator := startDT.Month()

	// Iterate through all the days between the start to end date.
	for _, todayDT := range dts {

		// Get weekday integer from the current iteration date.
		todayWeekdayInt := int(todayDT.Weekday())

		// If the picked day match the current day.
		if todayWeekdayInt == onFirstWeekday {

			// And if we are looking at the correct monthly then save.
			if monthIterator == todayDT.Month() {
				scheduleTimes = append(scheduleTimes, todayDT)

				// We already have the first occurance of the date for this
				// month so we can increment the month iterator so we know we
				// need to look for our next month.
				monthIterator++
			}
		}
	}

	return scheduleTimes
}

// GetDatesForLastWeekDayByMonthlyBasedRecurringSchedule Generates a list of datetimes based on a monthly recuring schedule which will find all the dates that fall on the weekday of the last week.
func GetDatesForLastWeekDayByMonthlyBasedRecurringSchedule(startDT time.Time, totalMonths int, onLastWeekday int) []time.Time {
	// Variable will store the datetimes that belong in the monthly based recurring schedule.
	scheduleTimes := []time.Time{}

	// Variable will calculate the last date based on total weeks in schedule.
	endDT := startDT.AddDate(0, totalMonths, 0) // Why "7"? Just to take into account extra week so we can handle for that particular months first week.

	// Create all the dates, incremented by day, from the start to finish datetimes.
	dts := RangeFromTimeStepper(startDT, endDT, 0, 0, 1, 0, 0, 0)

	monthIterator := startDT.Month()

	// Iterate through all the days between the start to end date.
	for _, todayDT := range dts {

		// Check to see if the current date we are iterating through is part of
		if IsTimeOnLastWeekOfMonth(todayDT) {

			// Get weekday integer from the current iteration date.
			todayWeekdayInt := int(todayDT.Weekday())

			// If the picked day match the current day.
			if todayWeekdayInt == onLastWeekday {

				// And if we are looking at the correct monthly then save.
				if monthIterator == todayDT.Month() {
					scheduleTimes = append(scheduleTimes, todayDT)

					// We already have the first occurance of the date for this
					// month so we can increment the month iterator so we know we
					// need to look for our next month.
					monthIterator++
				}
			}
		}

	}

	return scheduleTimes
}

// GetHourRange function will take a date value and return two date times:
// (1) The first date time will take the date and discard the minutes, so for
// example if you give 12:30 PM then it will return 12:00 PM.
// (2) The second date time will disacrd the minutes and increase by 1 hour,
// for example if you give 12:30 PM then it will return 1:00 PM.
// Therefore the purpose of this function is to provide an hourly range for
// the inputted parameter. For example if I say 12:300 PM then this function
// will return 12:00 PM & 1:00 AM.
//
// Developers Note: This would be useful to you if are building analytics engine
// which must group certain miunute based datapoints / timeseries into hours.
//
// Here are some more examples to help you visualize:
// Monday Dec 18th - 9:30 AM --> (1) Monday Dec 18th - 9:00 AM and (2) Monday Dec 18th - 10:00 AM
// Monday Dec 18th - 5:10 PM --> (1) Monday Dec 18th - 5:00 PM and (2) Monday Dec 18th - 7:00 PM
// Monday Dec 18th - 10:55 PM --> (1) Monday Dec 18th - 10:00 PM and (2) Monday Dec 18th - 11:00 PM
func GetHourRange(dt time.Time) (time.Time, time.Time) {
	// Discard the minutes and seconds to get the starting hour
	startHour := time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), 0, 0, 0, dt.Location())

	// Add 1 hour to get the ending hour while discarding the minutes
	endHour := startHour.Add(time.Hour)

	return startHour, endHour
}

// HourRangeForNow works just like the `GetHourRange` function however it works
// for the current date/time.
func HourRangeForNow(now func() time.Time) (time.Time, time.Time) {
	dt := now()
	return GetHourRange(dt)
}

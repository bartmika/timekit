package timekit

import (
	"time"
)

type TimeRange struct {
	Start time.Time
	End   time.Time
}

// HourlyRangeForTime function will take a date value and return two
// date times: (1) The first date time will take the date and discard the
// minutes, so for example if you give 12:30 PM then it will return 12:00 PM.
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
func HourlyRangeForTime(dt time.Time) *TimeRange {
	// Discard the minutes and seconds to get the starting hour
	startHour := time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), 0, 0, 0, dt.Location())

	// Add 1 hour to get the ending hour while discarding the minutes
	endHour := startHour.Add(time.Hour)

	return &TimeRange{
		Start: startHour,
		End:   endHour,
	}
}

// HourlyRangeForNow (TODO: Write detailed description)
func HourlyRangeForNow(now func() time.Time) *TimeRange {
	dt := now()
	return HourlyRangeForTime(dt)
}

// HourlyRangesBetweenTimes (TODO: Write detailed description)
func HourlyRangesBetweenTimes(start time.Time, end time.Time) []*TimeRange {
	dates := make([]*TimeRange, 0)

	// Developers Note:
	// We want to leverage our already unit tested code for the `range` functionality so we will use the `TimeStepper`
	// to iterate through the datetime values and add them to an `results` array.
	ts := NewTimeStepper(start, end, 0, 0, 0, 1, 0, 0)
	running := true
	for running {
		// Get the value we are on in the timestepper.
		v := ts.Get()

		dtr := HourlyRangeForTime(v)
		dates = append(dates, dtr)

		// Run our timestepper to get our next value.
		ts.Next()

		running = ts.Done() == false
	}

	return dates
}

// DailyRangeForTime (TODO: Write detailed description)
func DailyRangeForTime(dt time.Time) *TimeRange {
	startDay := time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
	endDay := startDay.Add(24 * time.Hour)

	return &TimeRange{
		Start: startDay,
		End:   endDay,
	}
}

// DailyRangeForNow (TODO: Write detailed description)
func DailyRangeForNow(now func() time.Time) *TimeRange {
	dt := now()
	return DailyRangeForTime(dt)
}

// DailyRangesBetweenTimes (TODO: Write detailed description)
func DailyRangesBetweenTimes(start time.Time, end time.Time) []*TimeRange {
	dates := make([]*TimeRange, 0)

	// Developers Note:
	// We want to leverage our already unit tested code for the `range` functionality so we will use the `TimeStepper`
	// to iterate through the datetime values and add them to an `results` array.
	ts := NewTimeStepper(start, end, 0, 0, 1, 0, 0, 0)
	running := true
	for running {
		// Get the value we are on in the timestepper.
		v := ts.Get()

		dtr := DailyRangeForTime(v)
		dates = append(dates, dtr)

		// Run our timestepper to get our next value.
		ts.Next()

		running = ts.Done() == false
	}

	return dates
}

// ISOWeeklyRangeForTime (TODO: Write detailed description)
func ISOWeeklyRangeForTime(dt time.Time) *TimeRange {
	dtFn := func() time.Time {
		return dt
	}
	startISOWeek := FirstDayOfThisISOWeek(dtFn)
	endISOWeek := LastDayOfThisISOWeek(dtFn)

	return &TimeRange{
		Start: startISOWeek,
		End:   endISOWeek,
	}
}

// ISOWeeklyRangeForNow (TODO: Write detailed description)
func ISOWeeklyRangeForNow(now func() time.Time) *TimeRange {
	dt := now()
	return ISOWeeklyRangeForTime(dt)
}

// ISOWeeklyRangesBetweenTimes (TODO: Write detailed description)
func ISOWeeklyRangesBetweenTimes(start time.Time, end time.Time) []*TimeRange {
	dates := make([]*TimeRange, 0)

	// Developers Note:
	// We want to leverage our already unit tested code for the `range` functionality so we will use the `TimeStepper`
	// to iterate through the datetime values and add them to an `results` array.
	ts := NewTimeStepper(start, end, 0, 0, 7, 0, 0, 0)
	running := true
	for running {
		// Get the value we are on in the timestepper.
		v := ts.Get()

		dtr := ISOWeeklyRangeForTime(v)
		dates = append(dates, dtr)

		// Run our timestepper to get our next value.
		ts.Next()

		running = ts.Done() == false
	}

	return dates
}

// MonthlyRangeForTime (TODO: Write detailed description)
func MonthlyRangeForTime(dt time.Time) *TimeRange {
	dtFn := func() time.Time {
		return dt
	}
	startMonth := FirstDayOfThisMonth(dtFn)
	endMonth := FirstDayOfNextMonth(dtFn)

	return &TimeRange{
		Start: startMonth,
		End:   endMonth,
	}
}

// MonthlyRangeForNow (TODO: Write detailed description)
func MonthlyRangeForNow(now func() time.Time) *TimeRange {
	dt := now()
	return MonthlyRangeForTime(dt)
}

// MonthlyRangesBetweenTimes (TODO: Write detailed description)
func MonthlyRangesBetweenTimes(start time.Time, end time.Time) []*TimeRange {
	dates := make([]*TimeRange, 0)

	// Developers Note:
	// We want to leverage our already unit tested code for the `range` functionality so we will use the `TimeStepper`
	// to iterate through the datetime values and add them to an `results` array.
	ts := NewTimeStepper(start, end, 0, 1, 0, 0, 0, 0)
	running := true
	for running {
		// Get the value we are on in the timestepper.
		v := ts.Get()

		dtr := MonthlyRangeForTime(v)
		dates = append(dates, dtr)

		// Run our timestepper to get our next value.
		ts.Next()

		running = ts.Done() == false
	}

	return dates
}

// YearlyRangeForTime (TODO: Write detailed description)
func YearlyRangeForTime(dt time.Time) *TimeRange {
	dtFn := func() time.Time {
		return dt
	}
	startYear := FirstDayOfThisYear(dtFn)
	endYear := FirstDayOfNextYear(dtFn)

	return &TimeRange{
		Start: startYear,
		End:   endYear,
	}
}

func YearlyRangeForNow(now func() time.Time) *TimeRange {
	dt := now()
	return YearlyRangeForTime(dt)
}

// YearlyRangesBetweenTimes (TODO: Write detailed description)
func YearlyRangesBetweenTimes(start time.Time, end time.Time) []*TimeRange {
	dates := make([]*TimeRange, 0)

	// Developers Note:
	// We want to leverage our already unit tested code for the `range` functionality so we will use the `TimeStepper`
	// to iterate through the datetime values and add them to an `results` array.
	ts := NewTimeStepper(start, end, 1, 0, 0, 0, 0, 0)
	running := true
	for running {
		// Get the value we are on in the timestepper.
		v := ts.Get()

		dtr := YearlyRangeForTime(v)
		dates = append(dates, dtr)

		// Run our timestepper to get our next value.
		ts.Next()

		running = ts.Done() == false
	}

	return dates
}

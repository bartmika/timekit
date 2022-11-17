package timekit

import "time"

// GetFutureDateByFiveMinuteIntervalPattern returns the future date that conforms to the 5 minute interval pattern. For example:
// Sunday Jan 9th - 1:00 AM --> Sunday Jan 9th - 1:00 AM
// Sunday Jan 9th - 1:01 AM --> Sunday Jan 9th - 1:05 AM
// Sunday Jan 9th - 1:28 AM --> Sunday Jan 9th - 1:30 AM
// Sunday Jan 9th - 1:59 AM --> Sunday Jan 9th - 2:00 AM
// Sunday Jan 9th - 11:59 PM --> Sunday Jan 10th - 12:00 AM
func GetFutureDateByFiveMinuteIntervalPattern(dt time.Time) time.Time {
	currMin := dt.Minute()
	availMins := [...]int{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55}

	var nextAvailH int = dt.Hour() + 1
	var nextAvailMin int
	for _, availMin := range availMins {
		if availMin >= currMin {
			nextAvailMin = availMin
			nextAvailH = dt.Hour() // Undue the next available hour.
			break
		}
	}
	return time.Date(dt.Year(), dt.Month(), dt.Day(), nextAvailH, nextAvailMin, 0, 0, dt.Location())
}

// GetFutureDateByTenMinuteIntervalPattern returns the future date that conforms to the 5 minute interval pattern. For example:
// Sunday Jan 9th - 1:00 AM --> Sunday Jan 9th - 1:00 AM
// Sunday Jan 9th - 1:01 AM --> Sunday Jan 9th - 1:10 AM
// Sunday Jan 9th - 1:28 AM --> Sunday Jan 9th - 1:30 AM
// Sunday Jan 9th - 1:59 AM --> Sunday Jan 9th - 2:00 AM
// Sunday Jan 9th - 11:59 PM --> Sunday Jan 10th - 12:00 AM
func GetFutureDateByTenMinuteIntervalPattern(dt time.Time) time.Time {
	currMin := dt.Minute()
	availMins := [...]int{0, 10, 20, 30, 40, 50}

	var nextAvailH int = dt.Hour() + 1
	var nextAvailMin int
	for _, availMin := range availMins {
		if availMin >= currMin {
			nextAvailMin = availMin
			nextAvailH = dt.Hour() // Undue the next available hour.
			break
		}
	}
	return time.Date(dt.Year(), dt.Month(), dt.Day(), nextAvailH, nextAvailMin, 0, 0, dt.Location())
}

// GetFutureDateByFiveteenMinuteIntervalPattern returns the future date that conforms to the 5 minute interval pattern. For example:
// Sunday Jan 9th - 1:00 AM --> Sunday Jan 9th - 1:00 AM
// Sunday Jan 9th - 1:01 AM --> Sunday Jan 9th - 1:15 AM
// Sunday Jan 9th - 1:28 AM --> Sunday Jan 9th - 1:30 AM
// Sunday Jan 9th - 1:59 AM --> Sunday Jan 9th - 2:00 AM
// Sunday Jan 9th - 11:59 PM --> Sunday Jan 10th - 12:00 AM
func GetFutureDateByFiveteenMinuteIntervalPattern(dt time.Time) time.Time {
	currMin := dt.Minute()
	availMins := [...]int{0, 15, 30, 45}

	var nextAvailH int = dt.Hour() + 1
	var nextAvailMin int
	for _, availMin := range availMins {
		if availMin >= currMin {
			nextAvailMin = availMin
			nextAvailH = dt.Hour() // Undue the next available hour.
			break
		}
	}
	return time.Date(dt.Year(), dt.Month(), dt.Day(), nextAvailH, nextAvailMin, 0, 0, dt.Location())
}

// GetFutureDateByThirtyMinuteIntervalPattern returns the future date that conforms to the 5 minute interval pattern. For example:
// Sunday Jan 9th - 1:00 AM --> Sunday Jan 9th - 1:00 AM
// Sunday Jan 9th - 1:01 AM --> Sunday Jan 9th - 1:30 AM
// Sunday Jan 9th - 1:28 AM --> Sunday Jan 9th - 1:30 AM
// Sunday Jan 9th - 1:59 AM --> Sunday Jan 9th - 2:00 AM
// Sunday Jan 9th - 11:59 PM --> Sunday Jan 10th - 12:00 AM
func GetFutureDateByThirtyMinuteIntervalPattern(dt time.Time) time.Time {
	currMin := dt.Minute()
	availMins := [...]int{0, 30}

	var nextAvailH int = dt.Hour() + 1
	var nextAvailMin int
	for _, availMin := range availMins {
		if availMin >= currMin {
			nextAvailMin = availMin
			nextAvailH = dt.Hour() // Undue the next available hour.
			break
		}
	}
	return time.Date(dt.Year(), dt.Month(), dt.Day(), nextAvailH, nextAvailMin, 0, 0, dt.Location())
}

// GetFutureDateByOneHourIntervalPattern returns the future date that conforms to the 5 minute interval pattern. For example:
// Sunday Jan 9th - 1:00 AM --> Sunday Jan 9th - 1:00 AM
// Sunday Jan 9th - 1:01 AM --> Sunday Jan 9th - 1:05 AM
// Sunday Jan 9th - 1:28 AM --> Sunday Jan 9th - 1:30 AM
// Sunday Jan 9th - 1:59 AM --> Sunday Jan 9th - 2:00 AM
// Sunday Jan 9th - 11:59 PM --> Sunday Jan 10th - 12:00 AM
func GetFutureDateByOneHourIntervalPattern(dt time.Time) time.Time {
	currHour := dt.Hour()
	currMinute := dt.Minute()

	if currMinute == 0 {
		return dt
	}
	return time.Date(dt.Year(), dt.Month(), dt.Day(), currHour+1, 0, 0, 0, dt.Location())
}

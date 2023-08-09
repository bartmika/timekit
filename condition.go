package timekit

import "time"

// IsMorning returns true if time is between 12 AM to 11:59 AM or (12 to 11:59) in 24 hour format
func IsMorning(t time.Time) bool {
	hour := t.Hour()
	return hour >= 0 && hour < 12
}

// IsAfternoon returns true if time is between 12PM to 4:59PM or (24 to 16:59) in 24 hour format
func IsAfternoon(t time.Time) bool {
	hour := t.Hour()
	return hour >= 12 && hour < 17
}

// IsEvening returns true if time is between 5PM to 7:59PM or (17 to 19:59) in 24 hour format
func IsEvening(t time.Time) bool {
	hour := t.Hour()
	return hour >= 17 && hour < 20
}

// IsNight returns true if time is between 8PM to 11:59 PM or (20 to 22:59) in 24 hour format
func IsNight(t time.Time) bool {
	hour := t.Hour()
	return hour >= 20 && hour < 24
}

// IsAfter6PM returns true if time is after 6PM.
func IsAfter6PM(t time.Time) bool {
	hour := t.Hour()
	return hour >= 18
}

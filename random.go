package timekit

import (
	"math/rand"
	"time"
)

// RandomDate Generates a random date between two given dates.
func RandomDate(startDate time.Time, endDate time.Time) time.Time {
	// The following code was inspired by: https://stackoverflow.com/a/43497333

	min := startDate.Unix()
	max := endDate.Unix()
	delta := max - min

	rand.Seed(time.Now().UnixNano()) // Set seed to increase randomness.
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

// DateInterval struct represents start and end dates.
type DateInterval struct {
	Start  time.Time
	Finish time.Time
}

// RandomDateIntervals Generates a list of date intervals between two dates.
func RandomDateIntervals(startDate time.Time, endDate time.Time, maxSeconds int64) []*DateInterval {
	var intervals []*DateInterval // Variable holds all the intervals we calculate.

	// Initialize variables for our calculations
	prevDate := startDate
	currDate := startDate
	randomDuration := time.Duration(rand.Int63n(maxSeconds)) * time.Second // Randomly generate duration (measured in seconds).
	nextDate := startDate.Add(randomDuration)
	safetyIter := 0 // Variable used to track number of loops and used to abort loop if too-many loops.

	// Keep performing our intervals calculation if (1) we haven't reached our
	// final date or (2) we looped too many times and possible programming error.
	for nextDate.Before(endDate) && safetyIter < 250_000 {
		interval := &DateInterval{
			Start:  currDate,
			Finish: nextDate,
		}
		intervals = append(intervals, interval)

		// Generate the next interval.
		randomDuration = time.Duration(rand.Int63n(maxSeconds)) * time.Second // Randomly generate new duration (measured in seconds).
		prevDate = currDate
		currDate = nextDate
		nextDate = currDate.Add(randomDuration)

		// For safety reasons keep a record of how many loops we've done.
		safetyIter++

		// If we are about to finish this current loop iteration and we detect
		// the finish date has been reached then generate our last record and
		// terminate the loop.
		if nextDate.After(endDate) {
			interval := &DateInterval{
				Start:  prevDate,
				Finish: endDate,
			}
			intervals = append(intervals, interval)
			break
		}
	}
	return intervals
}

// SegmentedDateInterval struct represents an date interval with a segment ID to track what segment the date interval belongs to.
type SegmentedDateInterval struct {
	ID       int64
	Interval *DateInterval
}

// RandomSegmentedDateIntervals Generates a segmented list of date intervals between two dates.
func RandomSegmentedDateIntervals(startDate time.Time, endDate time.Time, maxSeconds int64, totalSegments int64) []*SegmentedDateInterval {
	// Compute all our unsegmented intervals for our function.
	intervals := RandomDateIntervals(startDate, endDate, maxSeconds)

	// Reserve the memory.
	var segments = []*SegmentedDateInterval{}

	for _, interval := range intervals {
		randomIndex := rand.Int63n(totalSegments)

		segment := &SegmentedDateInterval{
			ID:       randomIndex,
			Interval: interval,
		}

		segments = append(segments, segment)
	}

	return segments
}

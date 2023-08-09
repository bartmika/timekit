package timekit

import (
	"testing"
	"time"
)

func TestRandomDate(t *testing.T) {
	startDateTime := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC) // Jan 1st 2022
	endDateTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)   // Jan 1st 2023
	result := RandomDate(startDateTime, endDateTime)

	// For debugging purposes only.
	// fmt.Println("result:", result)

	if result.Before(startDateTime) {
		t.Errorf("Incorrect date, got %s but was less then start date", result)
	}
	if result.After(endDateTime) {
		t.Errorf("Incorrect date, got %s but was greater then end date", result)
	}
}

func TestRandomDateIntervals(t *testing.T) {
	startDateTime := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC) // Jan 1st 2022
	endDateTime := time.Date(2022, 1, 1, 2, 0, 0, 0, time.UTC)   // Jan 1st 2023
	maxSeconds := int64(60)                                      // 1 minute
	drIntervals := RandomDateIntervals(startDateTime, endDateTime, maxSeconds)

	for _, dr := range drIntervals {
		if dr.Start.Before(startDateTime) {
			t.Errorf("Incorrect date, got %s but was less then start date", dr.Start)
		}
		if dr.Finish.After(endDateTime) {
			t.Errorf("Incorrect date, got %s but was greater then end date", dr.Finish)
		}

		// // For debugging purposes only.
		// fmt.Println("startDT:", dr.Start, "finishDT:", dr.Finish)
	}
}

func TestRandomSegmentedDateIntervals(t *testing.T) {
	startDateTime := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC) // Jan 1st 2022
	endDateTime := time.Date(2022, 1, 1, 2, 0, 0, 0, time.UTC)   // Jan 1st 2023
	maxSeconds := int64(60)                                      // 1 minute
	totalSegments := int64(3)
	segments := RandomSegmentedDateIntervals(startDateTime, endDateTime, maxSeconds, totalSegments)

	for _, segment := range segments {
		if segment.Interval.Start.Before(startDateTime) {
			t.Errorf("Incorrect date, got %s but was less then start date", segment.Interval.Start)
		}
		if segment.Interval.Finish.After(endDateTime) {
			t.Errorf("Incorrect date, got %s but was greater then end date", segment.Interval.Finish)
		}

		// // For debugging purposes only.
		// fmt.Printf("segment %v > interval %v\n", segment.ID, segment.Interval)
	}
}

package timekit

import (
	"testing"
	"time"
)

func TestEqualWithDrift(t *testing.T) {
	// DEVELOPERS NOTE:
	// On 2022-12-20 a poster on reddit named `_blue__sky_` asked the following
	// time-related question:
	//
	//     - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	//     I want to compare two time values but I want to allow (say) 2 seconds of difference as equal.
	//     For example
	//     2022-12-20T08:36:38Z and 2022-12-20T08:36:39Z should be equal while comparing.
	//
	//     Source:
	//     U/_blue__sky_. Comparing time values with margin of error. Reddit, 20 Dec. 2022, www.reddit.com/r/golang/comments/zqkegy/comparing_time_values_with_margin_of_error/
	//     - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	//
	// This is an excellent question and hence this library had `EqualWithDrift`
	// implement with the folowing unit test included as per the reddit users
	// post.

	t1, _ := ParseISO8601String("2022-12-20T08:36:38Z")
	t2, _ := ParseISO8601String("2022-12-20T08:36:39Z")
	errMargin := 2 * time.Second
	actual := EqualWithDrift(t1, t2, errMargin)
	if actual != true {
		t.Errorf("Incorrect result, got %v but was expecting %v", actual, true)
	}

	// If we lower the margin of error (drift) we should get false result.
	actual = EqualWithDrift(t1, t2, 2*time.Nanosecond)
	if actual != false {
		t.Errorf("Incorrect result, got %v but was expecting %v", actual, false)
	}
}

package timekit

import "time"

// EqualWithDrift function will compare two date/times but will allow duration margin of error (which we call drift) between the comparison. For example, if we have +/- 2 second margin of error (aka drift) and input 2022-12-20T08:36:38Z and 2022-12-20T08:36:39Z then the result should be equal.
func EqualWithDrift(t1 time.Time, t2 time.Time, drift time.Duration) bool {
	// DEVELOPERS NOTE:
	// The following code was copied from reddit. [0]
	//
	// [0]:
	// U/pdffs. Comparing time values with margin of error. Reddit, 20 Dec. 2022, www.reddit.com/r/golang/comments/zqkegy/comparing_time_values_with_margin_of_error/j0yif57/?context=3
	diff := t1.Sub(t2)
	return diff <= drift && diff >= -drift
}

package timekit

import (
	"time"
)

// monthAbbreviations is a mapping of time.Month values to their respective 3-letter
// abbreviations. This map is used to convert a full month name to its abbreviation.
var monthAbbreviations = map[time.Month]string{
	time.January:   "Jan",
	time.February:  "Feb",
	time.March:     "Mar",
	time.April:     "Apr",
	time.May:       "May",
	time.June:      "Jun",
	time.July:      "Jul",
	time.August:    "Aug",
	time.September: "Sep",
	time.October:   "Oct",
	time.November:  "Nov",
	time.December:  "Dec",
}

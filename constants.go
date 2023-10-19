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

// monthNumberAbbreviations is a mapping from month numbers (1 to 12) to their three-letter abbreviations.
var monthNumberAbbreviations = map[int]string{
	1:  "Jan", // January
	2:  "Feb", // February
	3:  "Mar", // March
	4:  "Apr", // April
	5:  "May", // May
	6:  "Jun", // June
	7:  "Jul", // July
	8:  "Aug", // August
	9:  "Sep", // September
	10: "Oct", // October
	11: "Nov", // November
	12: "Dec", // December
}

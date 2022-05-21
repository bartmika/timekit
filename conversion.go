package timekit

import (
	"strconv"
	"time"
)

// ParseJavaScriptTime will convert the number of milliseconds since the Unix Epoch parameter into Golang `time` format. As a result, the output of the JavaScript `getTime()` function can be used as the parameter in this function.
func ParseJavaScriptTime(i int64) time.Time {
	// Special thanks JavaScript timestamp to golang time.Time
	// https://gist.github.com/alextanhongpin/3b6b2ee47665ac9c1c32c805b86380a6
	return time.Unix(i/1000, (i%1000)*1000*1000)
}

// ParseJavaScriptTimeString will convert the string of milliseconds integers since the Unix Epoch parameter into Golang `time` format. As a result, the output of the JavaScript `getTime()` function can be used as the parameter in this function.
func ParseJavaScriptTimeString(s string) (time.Time, error) {
	i, valErr := strconv.ParseInt(s, 10, 64)
	if valErr != nil {
		return time.Now(), valErr
	}

	// Special thanks JavaScript timestamp to golang time.Time
	// https://gist.github.com/alextanhongpin/3b6b2ee47665ac9c1c32c805b86380a6
	return time.Unix(i/1000, (i%1000)*1000*1000), nil
}

// ToJavaScriptTime will return a Unix Epoch time value that your JavaScript code can read into JavaScript `Date` format. Example JavaScript code snippet of using the results of this function: `var date = new Date(UNIX_Timestamp * 1000);` as an example.
func ToJavaScriptTime(t time.Time) int64 {
	return t.Unix()
}

// ToISO8601String will convert the Golang `Date` format into an ISO 8601 formatted date/time string.
func ToISO8601String(t time.Time) string {
	return t.Format(time.RFC3339) // "How to convert ISO 8601 time in golang?" via https://stackoverflow.com/a/42217963
}

// ParseBubbleTime will convert the date/time string (ex: "Nov 11, 2011 11:00 am") used "https://bubble.io" into Golang `time`. You will find need of this function if the Bubble.io app you built will be making an API call to your Golang backend server.
func ParseBubbleTime(s string) (time.Time, error) {
	return time.Parse("Jan _2, 2006 15:04 am", s)
}

package timekit

import (
	"time"
)

// TimeFromJavaScriptGetTime will convert the number of milliseconds since the Unix Epoch parameter into Golang `time` format. As a result, the output of the JavaScript `getTime()` function can be used as the parameter in this function.
func TimeFromJavaScriptGetTime(i int64) time.Time {
	return time.Unix(i/1000, (i%1000)*1000*1000)
}

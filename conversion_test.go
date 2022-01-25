package timekit

import (
	"testing"
	"time"
)

func TestTimeFromJavaScriptGetTime(t *testing.T) {
	// EXAMPLE JAVASCRIPT CODE
	// var dt = new Date()
	// console.log(dt.getTime()) // 1643082322380
	// console.log(dt) // Mon Jan 24 2022 22:45:22 GMT-0500 (Eastern Standard Time)

	sample := int64(1643082322380)
	actual := TimeFromJavaScriptGetTime(sample)
	loc, _ := time.LoadLocation("America/Toronto")
	expected := time.Date(2022, 1, 24, 22, 45, 22, 380000000, loc) // 2022-01-24 22:45:22.38 -0500 EST
	if actual.Equal(expected) == false {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

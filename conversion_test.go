package timekit

import (
	"testing"
	"time"
)

func TestParseJavaScriptTime(t *testing.T) {
	// EXAMPLE JAVASCRIPT CODE
	// var dt = new Date()
	// console.log(dt.getTime()) // 1643082322380
	// console.log(dt) // Mon Jan 24 2022 22:45:22 GMT-0500 (Eastern Standard Time)

	jsTime := int64(1643082322380)
	actual := ParseJavaScriptTime(jsTime)
	loc, _ := time.LoadLocation("America/Toronto")
	expected := time.Date(2022, 1, 24, 22, 45, 22, 380000000, loc) // 2022-01-24 22:45:22.38 -0500 EST
	if actual.Equal(expected) == false {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}
}

func TestParseJavaScriptTimeString(t *testing.T) {
	// EXAMPLE JAVASCRIPT CODE
	// var dt = new Date()
	// console.log(dt.getTime()) // 1643082322380
	// console.log(dt) // Mon Jan 24 2022 22:45:22 GMT-0500 (Eastern Standard Time)

	jsTimeStr := "1643082322380"
	actual, err := ParseJavaScriptTimeString(jsTimeStr)
	if err != nil {
		t.Error(err)
	}

	// Case 1 of 2 - Correct string.

	loc, _ := time.LoadLocation("America/Toronto")
	expected := time.Date(2022, 1, 24, 22, 45, 22, 380000000, loc) // 2022-01-24 22:45:22.38 -0500 EST
	if actual.Equal(expected) == false {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}

	// Case 2 of 2 - Incorrect string.

	_, err = ParseJavaScriptTimeString("-------")
	if err == nil {
		t.Error("Incorrect error, should be not be nil but is nil!")
	}
}

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

func TestToJavaScriptTime(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")
	goTime := time.Date(2022, 1, 24, 22, 45, 22, 380000000, loc) // 2022-01-24 22:45:22.38 -0500 EST

	actual := ToJavaScriptTime(goTime)
	expected := int64(1643082322)

	if actual != expected {
		t.Errorf("Incorrect JavaScript UNIX date, got %d but was expecting %d", actual, expected)
	}

	// DEVELOPERS NOTE:
	// To confirm the above successfully works, open up your favourit web-browser and open up the inspector panel.
	// In your console, copy and paste the following and verify the results:
	//     var UNIX_Timestamp = 1643082322;
	//     var date = new Date(UNIX_Timestamp * 1000);
	//     console.log(date); // OUTPUT: Mon Jan 24 2022 22:45:22 GMT-0500 (Eastern Standard Time)
}

func TestToISO8601String(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")
	goTime := time.Date(2022, 1, 24, 22, 45, 22, 380000000, loc) // 2022-01-24 22:45:22.38 -0500 EST

	actual := ToISO8601String(goTime)
	expected := "2022-01-24T22:45:22-05:00"

	if actual != expected {
		t.Errorf("Incorrect ISO 8601 date, got %v but was expecting %v", actual, expected)
	}
}

func TestParseBubbleTime(t *testing.T) {
	// EXAMPLE JAVASCRIPT CODE
	// var dt = new Date()
	// console.log(dt.getTime()) // 1643082322380
	// console.log(dt) // Mon Jan 24 2022 22:45:22 GMT-0500 (Eastern Standard Time)

	bubbleTimeString := "Nov 11, 2011 11:00 am"
	actual, err := ParseBubbleTime(bubbleTimeString)
	if err != nil {
		t.Error(err)
	}

	// Case 1 of 2 - Correct string.

	loc, _ := time.LoadLocation("UTC")
	expected := time.Date(2011, 11, 11, 11, 0, 0, 000000000, loc)
	if actual.Equal(expected) == false {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}

	// Case 2 of 2 - Incorrect string.

	_, err = ParseJavaScriptTimeString("-------")
	if err == nil {
		t.Error("Incorrect error, should be not be nil but is nil!")
	}
}

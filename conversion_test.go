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

func TestParseISO8601String(t *testing.T) {

	actual, err := ParseISO8601String("2021-10-25T09:25:00.000")
	if err != nil {
		t.Error(err)
	}

	// Case 1 of 3 - Correct string (see: https://stackoverflow.com/a/38596248)

	loc, _ := time.LoadLocation("UTC")
	expected := time.Date(2021, 10, 25, 9, 25, 0, 000000000, loc)
	if actual.Equal(expected) == false {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}

	// Case 2 of 3 - Correct string (see: https://stackoverflow.com/a/38596248)

	actual, err = ParseISO8601String("2021-10-25T02:22:33+0000")
	if err != nil {
		t.Error(err)
	}

	expected = time.Date(2021, 10, 25, 2, 22, 33, 000000000, loc)
	if actual.Equal(expected) == false {
		t.Errorf("Incorrect date, got %s but was expecting %s", actual, expected)
	}

	// Case 3 of 3 - Incorrect string.

	_, err = ParseJavaScriptTimeString("-------")
	if err == nil {
		t.Error("Incorrect error, should be not be nil but is nil!")
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

func TestParseHourMinuteSecondDurationString(t *testing.T) {
	expected := map[string]time.Duration{
		"00:00:00": time.Duration(0),
		"00:00:01": time.Duration(1000 * 1000 * 1000),                // 1000 million nanoseconds is 1000 milliseconds is 1 second.
		"00:00:10": time.Duration(1000 * 1000 * 1000 * 10),           // 10 seconds.
		"00:01:00": time.Duration(1000 * 1000 * 1000 * 60),           // 1 minute.
		"00:10:00": time.Duration(1000 * 1000 * 1000 * 60 * 10),      // 10 minute.
		"01:00:00": time.Duration(1000 * 1000 * 1000 * 60 * 60),      // 1 hour.
		"10:00:00": time.Duration(1000 * 1000 * 1000 * 60 * 60 * 10), // 10 hours.
	}

	// Case 1 of 2 - Correct string.

	for k, v := range expected {
		actual, err := ParseHourMinuteSecondDurationString(k)
		if err != nil {
			t.Error(err)
		}
		if actual != v {
			t.Errorf("Incorrect value, got %s but was expecting %s for %s", actual, v, k)
		}
	}

	//
	// // Case 2 of 2 - Incorrect string.

	_, err := ParseHourMinuteSecondDurationString("-------")
	if err == nil {
		t.Error("Incorrect error, should be not be nil but is nil!")
	}
}

func TestToAmericanDateTimeString(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")
	goTime := time.Date(2022, 1, 24, 22, 45, 22, 380000000, loc) // 2022-01-24 22:45:22.38 -0500 EST

	actual := ToAmericanDateTimeString(goTime)
	expected := "January 24, 2022 10:45:22 PM"

	if actual != expected {
		t.Errorf("Incorrect ISO 8601 date, got %v but was expecting %v", actual, expected)
	}
}

func TestToAmericanDateString(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")
	goTime := time.Date(2022, 1, 24, 22, 45, 22, 380000000, loc) // 2022-01-24 22:45:22.38 -0500 EST

	actual := ToAmericanDateString(goTime)
	expected := "January 24, 2022"

	if actual != expected {
		t.Errorf("Incorrect ISO 8601 date, got %v but was expecting %v", actual, expected)
	}
}

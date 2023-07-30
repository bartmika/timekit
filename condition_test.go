package timekit

import (
	"testing"
	"time"
)

func TestIsMorning(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	if IsMorning(time.Date(2022, 1, 24, 22, 45, 22, 380000000, loc)) { // 2022-01-24 22:45:22.38 -0500 EST
		t.Error("Incorrect condition, expected false but got true")
	}
	if !IsMorning(time.Date(2022, 1, 24, 1, 45, 22, 380000000, loc)) { // 2022-01-24 1:45:22.38 -0500 EST
		t.Error("Incorrect condition, expected true but got false")
	}
}

func TestIsAfternoon(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	if IsAfternoon(time.Date(2022, 1, 24, 22, 45, 22, 380000000, loc)) { // 2022-01-24 22:45:22.38 -0500 EST
		t.Error("Incorrect condition, expected false but got true")
	}
	if !IsAfternoon(time.Date(2022, 1, 24, 13, 45, 22, 380000000, loc)) { // 2022-01-24 13:45:22.38 -0500 EST
		t.Error("Incorrect condition, expected true but got false")
	}
}

func TestIsEvening(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	if IsEvening(time.Date(2022, 1, 24, 22, 45, 22, 380000000, loc)) { // 2022-01-24 22:45:22.38 -0500 EST
		t.Error("Incorrect condition, expected false but got true")
	}
	if !IsEvening(time.Date(2022, 1, 24, 18, 45, 22, 380000000, loc)) { // 2022-01-24 18:45:22.38 -0500 EST
		t.Error("Incorrect condition, expected true but got false")
	}
}

func TestIsNight(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	if IsNight(time.Date(2022, 1, 24, 1, 45, 22, 380000000, loc)) { // 2022-01-24 1:45:22.38 -0500 EST
		t.Error("Incorrect condition, expected false but got true")
	}
	if !IsNight(time.Date(2022, 1, 24, 22, 45, 22, 380000000, loc)) { // 2022-01-24 22:45:22.38 -0500 EST
		t.Error("Incorrect condition, expected true but got false")
	}
}

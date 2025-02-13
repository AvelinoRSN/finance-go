package util

import "testing"

func TestStringTotime(t *testing.T) {
	var convertedTime = StringTotime("2019-02-12T09:00:00")

	if convertedTime.Year() != 2019 {
		t.Error("Year not equal")
	}
	if convertedTime.Month() != 2 {
		t.Error("Month not equal")
	}
	if convertedTime.Day() != 12 {
		t.Error("Day not equal")
	}
}
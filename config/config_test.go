package config

import "testing"

func TestTimeBetween(t *testing.T) {
	var between string = "20:00-23:40"
	from, to := parseTimeBetween(between)

	if from != 2000 {
		t.Error("from time error")
	}
	if to != 2340 {
		t.Error("to time error")
	}
}

func TestTimeBetweenToMorning(t *testing.T) {
	var between string = "20:00-07:00"
	from, to := parseTimeBetween(between)

	if from != 2000 {
		t.Error("from time error")
	}
	if to != 700 {
		t.Error("to time error")
	}
}

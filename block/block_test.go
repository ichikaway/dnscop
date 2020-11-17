package block

import (
	"fmt"
	"testing"
)

func TestTimeBetweenIn(t *testing.T) {
	from := 2000
	to := 2300
	now := 2010

	if !between(from, to, now) {
		s := fmt.Sprintf("from: %d, to: %d, now: %d", from, to, now)
		t.Error("expect not between. " + s)
	}
}

func TestTimeBetweenOut1(t *testing.T) {
	from := 2000
	to := 2300
	now := 1900

	if between(from, to, now) {
		s := fmt.Sprintf("from: %d, to: %d, now: %d", from, to, now)
		t.Error("expect not between. " + s)
	}
}

func TestTimeBetweenOut2(t *testing.T) {
	from := 2000
	to := 2300
	now := 2310

	if between(from, to, now) {
		s := fmt.Sprintf("from: %d, to: %d, now: %d", from, to, now)
		t.Error("expect not between. " + s)
	}
}

func TestTimeBetweenInToMorning(t *testing.T) {
	from := 2000
	to := 700
	now := 2010

	if !between(from, to, now) {
		s := fmt.Sprintf("from: %d, to: %d, now: %d", from, to, now)
		t.Error("expect not between. " + s)
	}
}
func TestTimeBetweenInToMorning2(t *testing.T) {
	from := 2000
	to := 700
	now := 610

	if !between(from, to, now) {
		s := fmt.Sprintf("from: %d, to: %d, now: %d", from, to, now)
		t.Error("expect not between. " + s)
	}
}

func TestTimeBetweenOutToMorning1(t *testing.T) {
	from := 2000
	to := 700
	now := 1900

	if between(from, to, now) {
		s := fmt.Sprintf("from: %d, to: %d, now: %d", from, to, now)
		t.Error("expect not between. " + s)
	}
}

func TestTimeBetweenOutToMorning2(t *testing.T) {
	from := 2000
	to := 700
	now := 810

	if between(from, to, now) {
		s := fmt.Sprintf("from: %d, to: %d, now: %d", from, to, now)
		t.Error("expect not between. " + s)
	}
}

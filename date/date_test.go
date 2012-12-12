// copyright 2012 Jacob Pipkin
//
// DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
// TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
// 0. You just DO WHAT THE FUCK YOU WANT TO.
package date

import (
	"time"
	"testing"
)

var n = time.Now()
var dateTime = time.Date(1979, time.September, 10, 22, 05, 01, 01, n.Location())

func TestMDY01(t *testing.T) {
	mdy := MDY(dateTime)
	if mdy != "09101979" {
		t.Error("MDY(dateTime), expected 09101979, got "+ mdy)
	}
}

func TestMDYs01(t *testing.T) {
	mdys := MDYs(dateTime, "-")
	if mdys != "09-10-1979" {
		t.Error("MDYs(dateTime, \"-\"), expected 09-10-1979, got "+ mdys)
	}
}

func TestHHMM01(t *testing.T) {
	hhmm := HHMM(dateTime)
	if hhmm != "2205" {
		t.Error("HHMM(dateTime), expected 2205, got "+ hhmm)
	}
}

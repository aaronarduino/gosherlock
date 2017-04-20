package gosherlock

import (
	"testing"
	"time"
)

var debug = false

func TestParse(t *testing.T) {
	// "February 24 at 3pm - 2pm March 3"
	input := "February 24 at 3pm - 2pm March 3"
	test := Parser{Config{DisableRanges: false}}
	testOutput := test.Parse(input)
	for _, s := range testOutput.Tokens {
		t.Log(s + "|")
	}
	if len(testOutput.Tokens) != 2 || debug {
		t.Fail()
	}
}

func TestIsSameDay(t *testing.T) {
	date1 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	date2 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	if IsSameDay(date1, date2) == false {
		t.Fail()
	}
}

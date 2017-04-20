package gosherlock

import "testing"

func TestParse(t *testing.T) {
	// "February 24 at 3pm - 2pm March 3"
	input := "February 24 at 3pm - 2pm March 3"
	test := Parser{Config{DisableRanges: false}}
	testOutput := test.Parse(input)
	for _, s := range testOutput.Tokens {
		t.Log(s + "|")
	}
	t.Fail()
}

package gosherlock

import (
	"regexp"
	"strings"
)

const rangeSplitters = `(\bto\b|\-|\b(?:un)?till?\b|\bthrough\b|\bthru\b|\band\b|\bends?\b)`

type Parser struct {
	Config
}

type Config struct {
	DisableRanges bool
}

type SherlockDate struct {
	Result string
	Str    string
	Ret    string
	Tokens []string
}

// Parse takes a string representing some English phrase, and
// returns an Event struct.
func (p *Parser) Parse(input string) []string {
	// tokenize the string
	var rangeSplit = regexp.MustCompile(rangeSplitters)
	date := SherlockDate{}

	date.Tokens = rangeSplit.Split(strings.ToLower(input), -1)
	return date
}

// parser is the main date parsing func
func parser(str, time, startTime string) {
	//
}

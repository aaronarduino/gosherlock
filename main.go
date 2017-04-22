package gosherlock

import (
	"regexp"
	"strings"
	"time"
)

const rangeSplitters = `(\bto\b|\-|\b(?:un)?till?\b|\bthrough\b|\bthru\b|\band\b|\bends?\b)`

type Parser struct {
	Config
	nowDate time.Time
}

type Config struct {
	DisableRanges bool
}

type ReturnedDate struct {
	IsAllDay   bool
	EventTitle string
	StartDate  time.Time
	EndDate    time.Time
}

type SherlockDate struct {
	Str    string
	Tokens []string
}

// Parse takes a string representing some English phrase, and
// returns a time.Time struct.
func (p *Parser) Parse(input string) SherlockDate {
	// tokenize the string
	var rangeSplit = regexp.MustCompile(rangeSplitters)
	date := SherlockDate{}

	date.Tokens = rangeSplit.Split(strings.ToLower(input), -1)
	date.Str = input

	return date
}

// parser is the main date parsing func
func (p *Parser) Parser(str, time, startTime string) {
	//
}

func (p *Parser) GetNow() time.Time {
	tst := time.Time{}
	if p.nowDate == tst {
		return p.nowDate
	}
	return time.Now()
}

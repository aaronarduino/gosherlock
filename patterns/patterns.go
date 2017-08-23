package patterns

import (
	"regexp"
	"strings"

	"github.com/aaronarduino/gosherlock/helpers"
)

const (
	RangeSplitters = `(\bto\b|\-|\b(?:un)?till?\b|\bthrough\b|\bthru\b|\band\b|\bends?\b)`
	EscapeRegex    = `[\-\[\]\/\{\}\(\)\*\+\?\.\\\^\$\|]`
	// oct, october
	Months = `\\b(jan(?:uary)?|feb(?:ruary)?|mar(?:ch)?|apr(?:il)?|may|jun(?:e)?|jul(?:y)?|aug(?:ust)?|sep(?:tember)?|oct(?:ober)?|nov(?:ember)?|dec(?:ember)?)\\b`
	// 3, 31, 31st, fifth
	Days = `\\b(?:(?:(?:on )?the )(?=\\d\\d?(?:st|nd|rd|th)))?([1-2]\\d|3[0-1]|0?[1-9])(?:st|nd|rd|th)?(?:,|\\b)`
	// 2014, 1990
	// Does not recognize 1930 for example because that could be confused with a valid time.
	// Exceptions are made for years in 21st century.
	Years = `\\b(20\\d{2}|\\d{2}[6-9]\\d)\\b`

	// 5/12/2014
	ShortForm = `\b(0?[1-9]|1[0-2])\/([1-2]\d|3[0-1]|0?[1-9])\/?(\d{2,4})?\b`

	// tue, tues, tuesday
	WeekdaysStr       = `\\b(sun|mon|tue(?:s)?|wed(?:nes)?|thu(?:rs?)?|fri|sat(?:ur)?)(?:day)?\\b`
	RelativeDateStr   = `((?:next|last|this) (?:week|month|year)|tom(?:orrow)?|tmrw|tod(?:ay)?|(?:right )?now|tonight|day after (?:tom(?:orrow)?|tmrw)|yest(?:erday)?|day before yest(?:erday)?)`
	InRelativeDateStr = `(\\d{1,4}|a) (day|week|month|year)s? ?(ago|old)?`

	InRelativeTime = `\b(\d{1,2} ?|a |an )(h(?:our)?|m(?:in(?:ute)?)?)s? ?(ago|old)?\b`
	InMilliTime    = `\b(\d+) ?(s(?:ec(?:ond)?)?|ms|millisecond)s? ?(ago|old)?\b`
	Midtime        = `(?:@ ?)?\b(?:at )?(dawn|morn(?:ing)?|noon|afternoon|evening|night|midnight)\b`
	// 23:50, 0700, 1900
	InternationalTime = `\b(?:(0[0-9]|1[3-9]|2[0-3]):?([0-5]\d))\b`
	// 5, 12pm, 5:00, 5:00pm, at 5pm, @3a
	ExplicitTime = `(?:@ ?)?\b(?:at |from )?(1[0-2]|[1-2]?[1-9])(?::?([0-5]\d))? ?([ap]\.?m?\.?)?(?:o'clock)?\b`

	More_than_comparator = `((?:more|greater|older|newer) than|after|before)`
	Less_than_comparator = `((?:less|fewer) than)`

	// filler words must be preceded with a space to count
	FillerWords = ` (from|is|was|at|on|for|in|due(?! date)|(?:un)?till?)\b`
	// less aggressive filler words regex to use when rangeSplitters are disabled
	FillerWords2 = ` (was|is|due(?! date))\b`
)

var (
	MonthDay                       = regexp.MustCompile(Months + ` ` + Days + `(?: ` + Years + `)?`)
	DayMonth                       = regexp.MustCompile(Days + `(?: (?:day )?of)? ` + Months + `(?: ` + Years + `)?`)
	DaysOnly                       = regexp.MustCompile(Days)
	Digit                          = regexp.MustCompile(`\\b(` + strings.Join(helpers.intToWords, "|") + `)\\b`)
	RelativeDate                   = regexp.MustCompile(`\\b` + RelativeDateStr + `\\b`)
	InRelativeDate                 = regexp.MustCompile(`\\b` + InRelativeDateStr + `\\b`)
	InRelativeDateFromRelativeDate = regexp.MustCompile(`\\b` + InRelativeDateStr + ` from ` + RelativeDateStr + `\\b`)
	Weekdays                       = regexp.MustCompile(`(?:(next|last) (?:week (?:on )?)?)?` + WeekdaysStr)
	OxtDays                        = regexp.MustCompile(`(?:\\boxt|\\bweek next) ` + WeekdaysStr)
	OxtDaysUK                      = regexp.MustCompile(WeekdaysStr + ` week\\b`)
)

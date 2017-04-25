package gosherlock

import (
	"regexp"
	"time"

	"github.com/aaronarduino/gosherlock/patterns"
)

func MonthDiff(d1, d2 time.Time) int {
	months := (d2.Year() - d1.Year()) * 12
	months -= int(d1.Month())
	months += int(d2.Month())
	if months <= 0 {
		return 0
	}
	return months
}

func EscapeRegExp(str string) string {
	tmp := regexp.MustCompile(patterns.EscapeRegex)
	return tmp.ReplaceAllLiteralString(str, `\$&`)
}

func IsSameDay(date1, date2 time.Time) bool {
	return date1.Month() == date2.Month()
}

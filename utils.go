package gosherlock

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aaronarduino/gosherlock/helpers"
	"github.com/aaronarduino/gosherlock/patterns"
)

func RelativeDateMatcher(match string, rtnTime *ReturnedDate) bool {
	switch match {
	case "next week":
		rtnTime.StartDate.AddDate(0, 0, 7)
	case "next month":
		rtnTime.StartDate.AddDate(0, 1, 0)
	case "next year":
		rtnTime.StartDate.AddDate(1, 0, 0)
	case "last week":
		rtnTime.StartDate.AddDate(0, 0, -7)
	case "last month":
		rtnTime.StartDate.AddDate(0, -1, 0)
	case "last year":
		rtnTime.StartDate.AddDate(-1, 0, 0)
	case "tom":
	case "tmrw":
	case "tomorrow":
		rtnTime.StartDate.AddDate(0, 0, 1)
	case "day after tom":
	case "day after tmrw":
	case "day after tomorrow":
		rtnTime.StartDate.AddDate(0, 0, 2)
	case "this week":
	case "this month":
	case "this year": // this week|month|year is pretty meaningless, but let's include it so that it parses as today
	case "tod":
	case "today":
		rtnTime.StartDate.AddDate(0, 0, 0)
	case "now":
	case "right now":
	case "tonight":
		rtnTime.StartDate.AddDate(0, 0, 0)
		// TODO: add case for tonight
	case "yest":
	case "yesterday":
		rtnTime.StartDate.AddDate(0, 0, -1)
	case "day before yest":
	case "day before yesterday":
		rtnTime.StartDate.AddDate(0, 0, -2)
	default:
		return false
	}
	rtnTime.HasYear = true
	return true
}

func InRelativeDateMatcher(num string, scale string, ago bool, rtnTime *ReturnedDate) bool {
	var number int

	// This may be a bad idea, but hey,
	// don't knock it till you try it, right?
	// This should match 'a' or 'an' or else parse number
	if num == "a" || num == "an" {
		number = 1
	} else {
		// in error-case, number should be zero and
		// number's null value is zero so err can be ignored
		number, _ := strconv.Atoi(num)
	}

	if ago {
		number = number * -1
	}

	switch scale {
	case "day":
		rtnTime.StartDate.AddDate(0, 0, number)
	case "week":
		rtnTime.StartDate.AddDate(0, 0, number*7)
	case "month":
		rtnTime.StartDate.AddDate(0, number, 0)
	case "year":
		rtnTime.StartDate.AddDate(number, 0, 0)
	default:
		return false
	}
	rtnTime.HasYear = true
	return true
}

func ChangeMonth(month string) int {
	return helpers.MonthToInt[month[0:3]]
}

func ChangeDay(inTime *time.Time, newDay int, hasNext string) {
	var diff = 7 - inTime.Day() + newDay
	if (diff > 7 && hasNext == "") || hasNext == "last" {
		diff -= 7
	} else if diff >= 0 && hasNext == "last" {
		diff -= 7
	} else if hasNext == "oxt" {
		diff += 7
	}
	*inTime.AddDate(0, 0, diff)
}

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

// StrToNum converts all the words in a string
// into numbers, such as four -> 4
func StrToNum(str string) string {
	tmp := regexp.MustCompile(``)
	return tmp.ReplaceAllStringFunc(str, func(val string) string {
		var out = strconv.Itoa(helpers.WordsToInt[val])
		if strings.Index(val[len(val)-2:], "th") != -1 {
			out += "th"
		} else if strings.Index(val[len(val)-2:], "st") != -1 {
			out += "st"
		} else if strings.Index(val[len(val)-2:], "nd") != -1 {
			out += "nd"
		} else if strings.Index(val[len(val)-2:], "rd") != -1 {
			out += "rd"
		}
		return out
	})

}

// NumToStr converts all the numbers in a string
// into regex for number|word, such as 4 -> 4|four
func NumToStr(str string) string {
	tmp := regexp.MustCompile(`((?:[1-9]|10)(?:st|nd|rd|th)?)`)
	return tmp.ReplaceAllStringFunc(str, func(val string) string {
		i, err := strconv.ParseInt("-42", 10, 64)
		if err != nil {
			log.Println(err)
			return ""
		}
		return "(?:" + val + "|" + intToWords[i-1] + ")"
	})
}

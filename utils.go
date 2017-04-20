package gosherlock

import "time"

func IsSameDay(date1, date2 time.Time) bool {
	return date1.Month() == date2.Month()
}

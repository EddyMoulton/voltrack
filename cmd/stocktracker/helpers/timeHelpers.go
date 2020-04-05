package helpers

import (
	"time"
)

func RemoveTime(input time.Time) time.Time {
	return time.Date(input.Year(), input.Month(), input.Day(), 0, 0, 0, 0, input.Location())
}

// rangeDate returns a date range function over start date to end date inclusive.
// After the end of the range, the range function returns a zero date,
// date.IsZero() is true.
func RangeDate(start, end time.Time) func() (time.Time, int) {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	index := 0

	inc := func() {
		index++
	}

	return func() (time.Time, int) {
		if start.After(end) {
			return time.Time{}, -1
		}
		date := start
		start = start.AddDate(0, 0, 1)
		defer inc()
		return date, index
	}
}

func DaysBetweenDatesInclusive(start, end time.Time) int {
	diff := end.Sub(start)

	return int(diff.Hours()/24) + 1
}

func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func OnSameDay(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}

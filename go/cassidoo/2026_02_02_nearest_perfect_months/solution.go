package nearestPerfectMonths

import (
	"fmt"
	"time"
)

func NearestPerfectMonths(year int) Solution {
	prevYearWithPerfectMonth := year
	nextYearWithPerfectMonth := year + 1

	for {
		if hasPerfectFebruary(prevYearWithPerfectMonth) {
			break
		}
		prevYearWithPerfectMonth--
	}
	for {
		if hasPerfectFebruary(nextYearWithPerfectMonth) {
			break
		}
		nextYearWithPerfectMonth++
	}

	return Solution{prev: fmt.Sprintf("%d-02", prevYearWithPerfectMonth), next: fmt.Sprintf("%d-02", nextYearWithPerfectMonth)}
}

func hasPerfectFebruary(year int) bool {
	firstIsMonday := time.Date(year, time.February, 1, 0, 0, 0, 0, time.UTC).Weekday() == time.Monday
	lastIsSunday := time.Date(year, time.February, 28, 0, 0, 0, 0, time.UTC).Weekday() == time.Sunday

	firstIsSunday := time.Date(year, time.February, 1, 0, 0, 0, 0, time.UTC).Weekday() == time.Sunday
	lastIsSaturday := time.Date(year, time.February, 28, 0, 0, 0, 0, time.UTC).Weekday() == time.Saturday

	isLeapyear := time.Date(year, time.February, 29, 0, 0, 0, 0, time.UTC).Month() == time.February
	return !isLeapyear && ((firstIsMonday && lastIsSunday) || (firstIsSunday && lastIsSaturday))
}

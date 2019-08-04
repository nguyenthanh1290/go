// Package leap provides a function that check whether a year is a leap year.
package leap

// IsLeapYear checks whether a year is a leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

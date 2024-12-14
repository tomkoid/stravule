package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// formatDate formats a date string into the desired format.
func formatDate(dateString string) string {
	date, err := ParseDate(dateString)
	if err != nil {
		fmt.Printf("Error parsing date: %v\n", err)
		return ""
	}

	weekday := date.Weekday().String()
	return fmt.Sprintf("%s %d. %d. %d", weekday, date.Day(), date.Month(), date.Year())
}

// for database timestamps
func NormalizeDateString(dateStr string) string {
	parts := strings.Split(dateStr, " ")
	if len(parts) != 4 {
		fmt.Println("normal part")
		return dateStr // Return as-is if unexpected format
	}

	timing := strings.Split(parts[1], ".")
	if len(timing) == 2 {
		timeParts := strings.Split(timing[1], ".")
		timing[1] = timeParts[0]
		parts[1] = timing[0]
	}

	return strings.Join(parts, " ")
}

// ParseDate parses a date string in the format "dd.mm.yyyy" and returns a time.Time object.
func ParseDate(dateString string) (time.Time, error) {
	parts := strings.Split(dateString, ".")
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("invalid date format: %s", dateString)
	}

	day, err := strconv.Atoi(parts[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid day: %s", parts[0])
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid month: %s", parts[1])
	}

	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid year: %s", parts[2])
	}

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	if date.IsZero() {
		return time.Time{}, fmt.Errorf("parsed date is invalid: %s", dateString)
	}

	return date, nil
}

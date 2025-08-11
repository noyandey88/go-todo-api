package utils

import "time"

// Epoch returns current time in seconds since Unix epoch
func Epoch() int64 {
	return time.Now().Unix()
}

// FormatDate returns time in human-readable format
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

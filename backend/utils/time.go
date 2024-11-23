package utils

import "time"

func GetTimestampString() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func GetTimestamp() time.Time {
	return time.Now().UTC()
}

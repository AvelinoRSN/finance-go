package util

import "time"

const layout = "2006-01-02T15:04:05"


func StringTotime(value string) time.Time {
	convertedTime, _ := time.Parse(layout, value)
	return convertedTime
}
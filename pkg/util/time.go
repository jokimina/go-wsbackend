package util

import "time"

const TimeFormat = "2006-01-02 15:04:05"

func GetHumanTimeNow() string {
	return time.Now().Format(TimeFormat)
}

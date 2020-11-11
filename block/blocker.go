package block

import (
	"regexp"
	"time"
)

/**
 * forbid watching youtube video later than 20:05
 */
func IsBlock(name string) bool {
	now := time.Now()
	nowInt := hourDateToInt(now.Hour(), now.Minute())
	limitHour := 20
	limitMin := 0
	limitInt := hourDateToInt(limitHour, limitMin)
	if nowInt > limitInt {
		r, _ := regexp.Compile("www.youtube.com|youtube.com|i.ytimg.com|.+.googlevideo.com")
		if r.MatchString(name) {
			return true
		}
	}
	return false
}

func hourDateToInt(hour int, min int) int {
	hour = hour * 100
	return hour + min
}

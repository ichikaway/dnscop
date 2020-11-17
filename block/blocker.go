package block

import (
	"dnscop/config"
	"regexp"
	"time"
)

/**
 * forbid watching youtube video later than 20:05
 */
func IsBlock(name string, conf *config.UserConfig) bool {
	now := time.Now()
	nowInt := hourDateToInt(now.Hour(), now.Minute())
	from := conf.GetFromTime()
	if from < nowInt {
		r, _ := regexp.Compile(conf.GetCondition())
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

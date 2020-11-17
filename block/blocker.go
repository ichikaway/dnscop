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
	to := conf.GetToTime()
	if between(from, to, nowInt) {
		r, _ := regexp.Compile(conf.GetCondition())
		if r.MatchString(name) {
			return true
		}
	}
	return false
}

func between(from int, to int, now int) bool {
	if from < now && now < to {
		return true
	}
	if to < from { //ex. to is morning time
		if from < now {
			return true
		}
		if now < to {
			return true
		}
	}
	return false
}

func hourDateToInt(hour int, min int) int {
	hour = hour * 100
	return hour + min
}

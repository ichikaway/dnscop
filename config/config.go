package config

import (
	"strconv"
	"strings"
)

type UserConfig struct {
	condition string
	timeFrom  int
	timeTo    int
}

func NewUserConfig(condition string, timeBetween string) *UserConfig {
	from, to := parseTimeBetween(timeBetween)
	return &UserConfig{
		condition: condition,
		timeFrom:  from,
		timeTo:    to,
	}
}

func (c UserConfig) GetCondition() string {
	return c.condition
}

func (c UserConfig) GetFromTime() int {
	return c.timeFrom
}

func (c UserConfig) GetToTime() int {
	return c.timeTo
}

func parseTimeBetween(timeBetween string) (int, int) {
	times := strings.Split(timeBetween, "-")
	from := strings.Split(times[0], ":")
	fromInt := hourDateToInt(from[0], from[1])
	to := strings.Split(times[1], ":")
	toInt := hourDateToInt(to[0], to[1])
	return fromInt, toInt
}

func hourDateToInt(h string, m string) int {
	hour, _ := strconv.Atoi(h)
	min, _ := strconv.Atoi(m)
	hour = hour * 100
	return hour + min
}

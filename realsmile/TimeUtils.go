package realsmile

import "time"

type TimeUtils struct {
}

var Time TimeUtils

const (
	template = "2006-01-02 15:04:05"
)

func (tu *TimeUtils) ParseBy(template, value string) (time.Time, error) {
	return time.Parse(template, value)
}
func (tu *TimeUtils) Parse(value string) (time.Time, error) {
	return time.Parse(template, value)
}
func (tu *TimeUtils) SnowFlakeStartTime(year int, month time.Month, day, hour, min, sec, nsec int) int64 {
	ti := time.Date(year, month, day, hour, min, sec, nsec, time.Local)
	return ti.UnixMilli()
}

package timeHelper

import (
	"strconv"
	"time"
)

func YearWeek(t time.Time) string {
	y, w := t.ISOWeek()
	return strconv.Itoa(y) + strconv.Itoa(w)
}

func GetNowYYMM() string {
	return time.Now().Format("0601")
}

func GetNowMMDD() string {
	return time.Now().Format("0102")
}

func GetNowYYYYMM() string {
	return time.Now().Format("200601")
}

func GetNowDate() string {
	return time.Now().Format("2006-01-02")
}

func GetDate(t time.Time) string {
	return t.Format("2006-01-02")
}

//获取几天前的时间
func GetPreDay(preDay int) time.Time {
	dis, _ := time.ParseDuration("-" + strconv.Itoa(24*preDay) + "h")
	return time.Now().Add(dis)
}

//获取几分钟后的时间
func GetNowAddMinTime(addMin int) time.Time {
	dis, _ := time.ParseDuration(strconv.Itoa(addMin) + "m")
	return time.Now().Add(dis)
}

//获取几分钟后的时间
func GetAddMinTime(t time.Time, addMin int) time.Time {
	dis, _ := time.ParseDuration(strconv.Itoa(addMin) + "m")
	return t.Add(dis)
}

//获取几秒后的时间
func GetNowAddSecondTime(t time.Time, addSecond int) time.Time {
	dis, _ := time.ParseDuration(strconv.Itoa(addSecond) + "s")
	return t.Add(dis)
}

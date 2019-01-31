// Format project Format.go
package Format

import (
	//	"log"
	"strconv"
	"strings"
	"time"
)

//日期时间格式串很奇葩，请注意格式串定义
const (
	//日期默认格式 yyyy-MM-dd
	DateFormatDefault string = "2006-01-02"
	//时间默认格式 HH:mm:ss
	TimeFormatDefault string = "15:04:05"
	//日期时间默认格式 yyyy-MM-dd HH:mm:ss
	DateTimeFormatDefault string = "2006-01-02 15:04:05"
	//日期时间默认格式 yyyyMMddHHmmss
	DateTimeFormat0 string = "20060102150405"
	//日期时间默认格式 yyyyMMddHHmm
	DateTimeFormat1 string = "200601021504"
	//日期时间默认格式 yyyy-MM-dd HH:mm:ss.ffffff 注意：原配置为2006-01-01 01:01:40.000000格式后输出错误
	TimeStampFormat0 string = "2006-01-02 15:04:05.000000"
	//日期时间默认格式 yyyyMMddHHmmssffffff
	TimeStampFormat1 string = "20060102150405000000"
	//日期时间默认格式 yyyyMMddHHmmss.ffffff
	TimeStampFormat2 string = "20060102150405.000000"
)

//获取日期时间格式串的分钟值
func getMinute(currTime string) int {
	mm := currTime[14:16]

	min, _ := strconv.Atoi(mm)

	return min
}

//时间加减
func addTimeResult(currTimeStr string, d time.Duration) string {
	var result string

	currTime, err := time.ParseInLocation(DateTimeFormatDefault, currTimeStr, time.Local)

	if err != nil {
		panic(err)
	}

	currTime = currTime.Add(d)

	result = currTime.Format(DateTimeFormatDefault)

	return result
}

//获取分钟
func GetMinuteTime(currTime string) (string, string) {
	var beginTimeStr, endTimeStr string

	beginTimeStr = strings.Join([]string{currTime[:17], "00"}, "")

	d, _ := time.ParseDuration("1m")
	endTimeStr = addTimeResult(beginTimeStr, d)

	return beginTimeStr, endTimeStr
}

//获取5分钟时间区间
//返回结果：开始时间和截止时间
//时间格式:yyyy-MM-dd HH:mm:ss
func GetFiveMinuteTime(currTime string) (string, string) {
	var beginTime, endTime string

	d, _ := time.ParseDuration("5m")

	min := getMinute(currTime)

	dataTimefrontMin := currTime[:14]
	dataTimebehindMin := currTime[16:]

	//2018-10-11 方法待优化
	//改进方法如下：
	//开始时间=计算分钟值拼接日期时间字符串，其分钟值=(分钟值- 分钟值除以5取余数)
	//截止时间=开始时间+5分钟

	remainder := min % 5

	minute := strconv.Itoa(min - remainder)

	if len(minute) == 1 {
		minute = strings.Join([]string{"0", minute}, "")
	}

	beginTime = strings.Join([]string{dataTimefrontMin, minute, dataTimebehindMin}, "")

	endTime = addTimeResult(beginTime, d)

	return beginTime, endTime
}

//获取15分钟时间区间
//返回结果：开始时间和截止时间
//时间格式:yyyy-MM-dd HH:mm:ss
func GetFifthMinuteTime(currTime string) (string, string) {
	var beginTime, endTime string

	d, _ := time.ParseDuration("15m")

	min := getMinute(currTime)

	dataTimefrontMin := currTime[:14]
	dataTimebehindMin := currTime[16:]

	//2018-10-11 方法待优化
	//改进方法如下：
	//开始时间=计算分钟值拼接日期时间字符串，其分钟值=(分钟值- 分钟值除以15取余数)
	//截止时间=开始时间+15分钟

	remainder := min % 15

	minute := strconv.Itoa(min - remainder)

	if len(minute) == 1 {
		minute = strings.Join([]string{"0", minute}, "")
	}

	beginTime = strings.Join([]string{dataTimefrontMin, minute, dataTimebehindMin}, "")

	endTime = addTimeResult(beginTime, d)

	return beginTime, endTime
}

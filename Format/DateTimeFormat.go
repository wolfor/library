// Format project Format.go
package Format

//日期时间格式串很奇葩，请注意格式串定义
const (
	//日期默认格式 yyyy-MM-dd
	DateFormatDefault string = "2006-01-02"
	//时间默认格式 HH:mm:ss
	TimeFormatDefault string = "15:04:05"
	//日期时间默认格式 yyyy-MM-dd HH:mm:ss
	DateTimeFormatDefault string = "2006-01-02 15:04:05"
	//日期时间默认格式 yyyy-MM-dd HH:mm:ss.ffffff 注意：原配置为2006-01-01 01:01:40.000000格式后输出错误
	TimeStampFormat0 string = "2006-01-02 15:04:05.000000"
)

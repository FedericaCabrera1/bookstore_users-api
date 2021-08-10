package date_utils

import "time"

const (
	apiDateLayout = "2019-01-02t15:04:052"
	apiDbLayout   = "2019-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC() //UTC to work in standard timezone
}
func GetNowString() string {
	return GetNow().Format(apiDateLayout) //take the local time and display it in this format (format in the const)
}
func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}

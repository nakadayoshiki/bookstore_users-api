package date

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// GetNow in japan
func GetNow() time.Time {
	timeLocation, _ := time.LoadLocation("Asia/Tokyo")
	return time.Now().In(timeLocation)
}

// GetNowString formats the time and date
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

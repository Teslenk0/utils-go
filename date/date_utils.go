package date


import "time"

const (
	apiDateLayout   = "2006-01-02T15:04:05Z"
	apiDbDateLayout = "2006-01-02 15:04:05"
)

//GetNow - Returns current time in standard zone (UTC)
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString - Same as GetNow but this returns as a formatted string
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDBString() string {
	return GetNow().Format(apiDbDateLayout)
}

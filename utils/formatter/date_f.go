package formatter

import "time"

func DateToDbDateFormat(data time.Time) string {
	return data.Format("2006-01-02")
}
func DateToShortFormat(data time.Time) string {
	return data.Format("01-02-2006")
}

func DateToMediumFormat(data time.Time) string {
	return data.Format("02-Jan-2006")
}

func DateToLongFormat(data time.Time) string {
	return data.Format("02 January 2006")
}

func TimeToLogFormat(data time.Time) string {
	return data.Format("20060102.150405")
}

func TimeToPeriod(data time.Time) string {
	return data.Format("20060102")
}

func TimeToMonth(data time.Time) string {
	return data.Format("1")
}

func TimeToMonth2(data time.Time) string {
	return data.Format("01")
}

func TimeToYear(data time.Time) string {
	return data.Format("2006")
}

func TimeToDbTimeFormat(data time.Time) string {
	return data.Format("2006-01-02 15:04:05")
}

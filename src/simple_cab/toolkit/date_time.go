package toolkit

import (
	"simple_cab/logging"
	"time"
)

var _ = logging.Info

func GetTimeFromDateTimeString(dateTimeStr string) (tm time.Time, err error) {
	return time.Parse(time.RFC3339, dateTimeStr)
}

func GetDiffByDays(startTime *time.Time, endTime *time.Time) int {
	timeDiff := endTime.Sub(*startTime)
	return int(timeDiff.Hours() / 24)
}

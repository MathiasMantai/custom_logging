package file_logging

import (
	"time"
)

// convert duration in days to nanoseconds
func DayToNano(days int64) int64 {
	return days * 24 * 3600 * 1000000
}

func GetCurrentTime(timeZone string, layout string) (string, error) {
    location, err := time.LoadLocation(timeZone)
    if err != nil {
        return "", err
    }

    // Convert the local time to German time
    germanTime := time.Now().In(location).Format(layout)
    return germanTime, nil
}
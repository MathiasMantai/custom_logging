package custom_logging

import (
	"time"
)

// convert duration in days to nanoseconds
func DayToNano(days int64) int64 {
	return days * 24 * 3600 * 1000000
}

func GetCurrentTime() (string, error) {
    germanLoc, err := time.LoadLocation("Europe/Berlin")
    if err != nil {
        return "", err
    }

	layout := "02.01.2006 15:04"

    // Convert the local time to German time
    germanTime := time.Now().In(germanLoc).Format(layout)
    return germanTime, nil
}
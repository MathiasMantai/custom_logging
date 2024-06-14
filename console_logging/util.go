package console_logging

import (
	"time"
)

func GetCurrentTime(timeZone string, layout string) (string, error) {
    location, err := time.LoadLocation(timeZone)
    if err != nil {
        return "", err
    }

    // Convert the local time to German time
    germanTime := time.Now().In(location).Format(layout)
    return germanTime, nil
}
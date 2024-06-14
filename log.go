package custom_logging

import (
	// "archive/zip"
	"fmt"
	"os"
	"strings"
	"time"
)

type Logger struct {
	KeepOldest int32
	DirPath    string
	LogPrefix  string
}

func (l *Logger) LogMsgs(msg string) bool {
	encodedMsg := []byte(msg)
	writeError := os.WriteFile(l.DirPath, encodedMsg, 0644)

	return writeError == nil
}

func (l *Logger) LogRotate() bool {
	files, readDirError := os.ReadDir(l.DirPath)

	if readDirError != nil {
		panic("x> error ")
	}

	for _, file := range files {
		fmt.Println(file.Name())
		logName := file.Name()
		logTime := strings.Replace(logName, l.LogPrefix, "", -1)
		tmpTime, parseTimeError := time.Parse("2024-12-01 13:25:00", logTime)

		if parseTimeError != nil {
			panic(parseTimeError)
		}

		since := time.Since(tmpTime)
		sinceInDays := (since.Hours() / 24)

		if sinceInDays >= float64(l.KeepOldest) {
			//zip log
		}
	}

	return true
}

func (l *Logger) InitLogDir() {
	createDirError := os.Mkdir(l.DirPath, 0755)
	if createDirError != nil {
		fmt.Println("=> Directory exists already. Skipping...")
	} else {
		fmt.Println("=> Log Directory created...")
	}
}

func CreateLogger(keppOldest int32, dirPath string, logPrefix string) Logger {
	if logPrefix == "" {
		logPrefix = "log_"
	}

	return Logger{
		keppOldest,
		dirPath,
		logPrefix,
	}
}

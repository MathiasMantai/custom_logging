package console_logging

type ConsoleLogger struct {
	TimeZone string	
}

//set the timezone for the logger. used for printing timestamps
func (cl *ConsoleLogger) SetTimeZone(timeZone string) {
	cl.TimeZone = timeZone
}
 

//create a new console logger

func NewConsoleLogger(timeZone string) *ConsoleLogger {
	return &ConsoleLogger{
		TimeZone: timeZone,
	}
}
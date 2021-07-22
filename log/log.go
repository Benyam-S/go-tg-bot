package log

// Debug is a constant that indicates the logger is in debug mode
const Debug = "Debug"

// Normal is a constant that indicates the logger is normally logging in the log files
const Normal = "Normal"

// None is a constant that indicates the logger isn't logging
const None = "None"

// BotLogFile is a constant that indicates the log will be written on 'BotLogFile'
const BotLogFile = "BotLogFile"

// ErrorLogFile is a constant that indicates the log will be written on 'ErrorLogFile'
const ErrorLogFile = "ErrorLogFile"

// ILogger is an interface that defines the logging style
type ILogger interface {
	SetFlag(state string)
	Log(stmt, logFile string)
	LogToParent(stmt string)
}

// LogContainer is a type that defines all the available logs
type LogContainer struct {
	BotLogFile   string // The 'BotLogFile' is used for logging any transaction that is done by the bot
	ErrorLogFile string // The 'ErrorLogFile' is used for logging any error occuring while performing transaction
}

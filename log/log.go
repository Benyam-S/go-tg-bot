package log

// ILogger is an interface that defines the logging style
type ILogger interface {
	SetFlag(state bool)
	Log(stmt, logFile string)
	LogToParent(stmt, stmtType string)
	LogFileError(stmt string)
	LogFileArchive(stmt string)
}

// LogContainer is a type that defines all the available logs
type LogContainer struct {
	ServerLogFile  string
	BotLogFile     string
	ErrorLogFile   string
	ArchiveLogFile string
}

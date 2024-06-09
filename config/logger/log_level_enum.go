package logger

type LogLevel string

const (
	LogLevel_Error   LogLevel = "Error"
	LogLevel_Warning LogLevel = "Warning"
	LogLevel_Info    LogLevel = "Info"
	LogLevel_Debug   LogLevel = "Debug"
	LogLevel_Success LogLevel = "Success"
)

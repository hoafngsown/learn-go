package logger

type LogType string

const (
	INFO  LogType = "INFO"
	DEBUG LogType = "DEBUG"
	WARN  LogType = "WARN"
	ERROR LogType = "ERROR"
)

type LogData map[string]interface{}

type LogUtil interface {
	Info(message string, data ...LogData)
	Debug(message string, data ...LogData)
	Warn(message string, data ...LogData)
	Error(message string, data ...LogData)
}

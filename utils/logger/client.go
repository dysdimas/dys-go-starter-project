package logger

type LogLevel uint32

const (
	FatalLevel LogLevel = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

type ClientLogger interface {
	Init(level LogLevel) error
	Debug(message string)
	DebugF(additionalData map[string]interface{}, message string, args ...interface{})
	Info(message string)
	InfoF(additionalData map[string]interface{}, message string, args ...interface{})
	Warn(message string)
	WarnF(additionalData map[string]interface{}, message string, args ...interface{})
	Error(message string)
	ErrorF(additionalData map[string]interface{}, message string, args ...interface{})
	Fatal(message string)
	FatalF(additionalData map[string]interface{}, message string, args ...interface{})
}

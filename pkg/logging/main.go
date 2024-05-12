package logging

type Logger interface {
	Error(s string)
	Info(s string)
	Warn(s string)
	Debug(s string)
}

type DefaultLogger struct{}

type NullLogger struct{}

func (nl NullLogger) Error(s string) {}
func (nl NullLogger) Info(s string)  {}
func (nl NullLogger) Warn(s string)  {}
func (nl NullLogger) Debug(s string) {}

func (l *DefaultLogger) Error(s string) {
	println("ERROR: " + s)
}

func (l *DefaultLogger) Info(s string) {
	println("INFO: " + s)
}

func (l *DefaultLogger) Warn(s string) {
	println("WARN: " + s)
}

func (l *DefaultLogger) Debug(s string) {
	println("DEBUG: " + s)
}

func NewLogger() Logger {
	return &DefaultLogger{}
}

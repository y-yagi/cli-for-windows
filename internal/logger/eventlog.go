// +build windows

package log

import "golang.org/x/sys/windows/svc/eventlog"

type Logger struct {
	eventlog *eventlog.Log
	app      string
}

func NewLogger(app string) (*Logger, error) {
	const supports = eventlog.Error | eventlog.Warning | eventlog.Info
	err := eventlog.InstallAsEventCreate(app, supports)
	if err != nil {
		return nil, err
	}

	l, err := eventlog.Open(app)
	if err != nil {
		eventlog.Remove(app)
		return nil, err
	}

	return &Logger{eventlog: l, app: app}, nil
}

func (logger *Logger) Close() {
	logger.eventlog.Close()
	eventlog.Remove(logger.app)
}

func (logger *Logger) Info(msg string) {
	logger.eventlog.Info(1, msg)
}

func (logger *Logger) Warning(msg string) {
	logger.eventlog.Warning(2, msg)
}

func (logger *Logger) Error(msg string) {
	logger.eventlog.Error(3, msg)
}

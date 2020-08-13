package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

var log = New()

type Logger struct {
	iLog *logrus.Logger
}

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

type Level uint32

func New() *Logger {
	return &Logger{
		iLog: logrus.New(),
	}
}

func SetLevel(level Level) {
	fmt.Println("set level", level)
	log.iLog.SetLevel(logrus.Level(level))
}

func Infoln(args ...interface{}) {
	log.iLog.Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	log.iLog.Infof(format, args...)
}

func Fatalln(args ...interface{}) {
	log.iLog.Fatalln(args...)
}

func Println(args ...interface{}) {
	log.iLog.Println(args...)
}

func Warnln(args ...interface{}) {
	log.iLog.Warnln(args...)
}

package logger

import (
	"fmt"
	"path"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

func New() logrus.FieldLogger {
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetLevel(logrus.TraceLevel)

	l.Formatter = newTextFormatter()

	return l
}

func newTextFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		FullTimestamp:    true,
		TimestampFormat:  time.RFC3339,
		CallerPrettyfier: callerPrettyfier,
	}
}

func callerPrettyfier(frame *runtime.Frame) (string, string) {
	return "", fmt.Sprintf("%s: %d", path.Base(frame.File), frame.Line)
}

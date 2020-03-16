package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

var (
	maxAge       = 7 * 24 * time.Hour
	rotationTime = 24 * time.Hour
)

// NewLogger ...
func NewLogger(logFile string, options ...time.Duration) {
	if len(options) == 1 {
		maxAge = options[0]
	}

	if len(options) == 2 {
		maxAge = options[0]
		rotationTime = options[1]
	}

	logrus.SetLevel(logrus.DebugLevel)

	lfshook := newLfsHook(logFile, maxAge, rotationTime)
	logrus.AddHook(lfshook)
}

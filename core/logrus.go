package core

import (
	"0049-server-go/global"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

// Colour
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// Format implement Formatter(entry *logrus.Entry) ([]byte, error) interface
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Display colours according to different levels
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	// If entry.Buffer is not nil, use it directly; otherwise, create a new one
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	log := global.Config.Logger // Get configuration from global variables

	// Custom date format
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	// HasCaller() Determine if the return function name and line number are enabled
	if entry.HasCaller() {
		// Custom file path
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line) // Obtain file name and line number
		// Custom Output Format
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", log.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", log.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	// return []byte
	return b.Bytes(), nil
}

// InitLogger is the function to initialize the logger.
func InitLogger() *logrus.Logger {
	mLog := logrus.New()                                // Create a new logrus instance
	mLog.SetOutput(os.Stdout)                           // Set Output Type
	mLog.SetReportCaller(global.Config.Logger.ShowLine) // Enable return name and line
	mLog.SetFormatter(&LogFormatter{})                  // Set Formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level)
	InitDefaultLogger() // InitDefaultLogger() Initialize global log
	return mLog
}

// InitDefaultLogger is the function to initialize the default logger.
func InitDefaultLogger() {
	// global log
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(global.Config.Logger.ShowLine)
	logrus.SetFormatter(&LogFormatter{})
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}

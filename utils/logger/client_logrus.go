package logger

import (
	"dys-go-starter-project/utils/formatter"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ClientLogrusConfig struct {
	MaxFileSizeMb  int
	MaxAgeDay      int
	Compress       bool
	Logrus         *logrus.Logger
	PrefixFileName string
	LogDirPath     string
}

type ClientLogrus struct {
	Config *ClientLogrusConfig
}

func (client *ClientLogrus) Init(level LogLevel) error {
	if client.Config.PrefixFileName == "" {
		client.Config.PrefixFileName = "logrus"
	}

	client.Config.Logrus = logrus.New()
	var logrusLevel logrus.Level

	switch level {
	case FatalLevel:
		logrusLevel = logrus.FatalLevel
	case ErrorLevel:
		logrusLevel = logrus.ErrorLevel
	case WarnLevel:
		logrusLevel = logrus.WarnLevel
	case InfoLevel:
		logrusLevel = logrus.InfoLevel
	case DebugLevel:
		logrusLevel = logrus.DebugLevel
	}

	client.Config.Logrus.Level = logrusLevel

	client.Config.Logrus.Out = &lumberjack.Logger{
		Filename: client.Config.LogDirPath + client.Config.PrefixFileName + "." + formatter.TimeToLogFormat(time.Now()) + "." + strconv.Itoa(os.Getpid()) + ".log",
		MaxSize:  client.Config.MaxFileSizeMb, // megabytes
		MaxAge:   client.Config.MaxAgeDay,     //days
		Compress: client.Config.Compress,      // disabled by default
	}
	return nil
}

func (client *ClientLogrus) Debug(message string) {
	client.Config.Logrus.Debug(message)
}

func (client *ClientLogrus) DebugF(additionalData map[string]interface{}, message string, args ...interface{}) {
	client.Config.Logrus.WithFields(additionalData).Debugf(message, args)
}

func (client *ClientLogrus) Info(message string) {
	client.Config.Logrus.Info(message)
}

func (client *ClientLogrus) InfoF(additionalData map[string]interface{}, message string, args ...interface{}) {
	client.Config.Logrus.WithFields(additionalData).Infof(message, args)
}

func (client *ClientLogrus) Warn(message string) {
	client.Config.Logrus.Warn(message)
}

func (client *ClientLogrus) WarnF(additionalData map[string]interface{}, message string, args ...interface{}) {
	client.Config.Logrus.WithFields(additionalData).Warnf(message, args)
}

func (client *ClientLogrus) Error(message string) {
	client.Config.Logrus.Error(message)
}

func (client *ClientLogrus) ErrorF(additionalData map[string]interface{}, message string, args ...interface{}) {
	client.Config.Logrus.WithFields(additionalData).Errorf(message, args)
}

func (client *ClientLogrus) Fatal(message string) {
	client.Config.Logrus.Fatal(message)
}

func (client *ClientLogrus) FatalF(additionalData map[string]interface{}, message string, args ...interface{}) {
	client.Config.Logrus.WithFields(additionalData).Fatalf(message, args)
}

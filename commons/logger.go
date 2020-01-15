package commons

import (
	"os"

	"go.uber.org/zap"
)

// ILogger represent common interface
type ILogger interface {
	Warn(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

// Logger is ILogger instance
var Logger ILogger

func init() {
	env := os.Getenv("env")
	if env == "" || env == "local" || env == "dev" || env == "development" {
		sugar, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		Logger = sugar.Sugar()
		return
	}
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"/home/work/logs/server.log"}
	config.ErrorOutputPaths = []string{"/home/work/logs/err.log"}
	sugar, err := config.Build()
	if err != nil {
		panic(err)
	}
	Logger = sugar.Sugar()
	return
}

// SetLogger init the Logger
func SetLogger(logger ILogger) {
	if logger != nil {
		Logger = logger
		return
	}
}

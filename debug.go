package alipaysdk

import (
	"log"
	"os"
)

var enableDebug = false

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[alipay sdk] ", 0)
}

func SetDebug(debug bool) {
	enableDebug = debug
}

func SetLogger(log *log.Logger) {
	if log != nil {
		logger = log
	}
}

func Printf(format string, v ...interface{}) {
	if enableDebug {
		logger.Printf(format, v...)
	}
}

package log

import (
	"github.com/soerjadi/stockist/internal/pkg/log/logger"
)

func InitLog(pathFile, appName string) {
	path := pathFile + appName

	_errorLogger.SetConfig(&logger.Config{
		Level:    logger.ErrorLevel,
		LogFile:  path + ".error.log",
		Caller:   true,
		UseColor: true,
		UseJSON:  true,
	})

	_infoLogger.SetConfig(&logger.Config{
		Level:    logger.InfoLevel,
		LogFile:  path + ".info.log",
		Caller:   false,
		UseColor: true,
	})

}

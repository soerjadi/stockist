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

	// _debugLogger.SetConfig(&logger.Config{
	// 	Level:    logger.DebugLevel,
	// 	LogFile:  path + ".debug.log",
	// 	Caller:   true,
	// 	UseColor: true,
	// 	UseJSON:  true,
	// })

	// _fatalLogger.SetConfig(&logger.Config{
	// 	Level:    logger.FatalLevel,
	// 	LogFile:  path + ".fatal.log",
	// 	Caller:   true,
	// 	UseColor: true,
	// 	UseJSON:  true,
	// })

	// _warnLogger.SetConfig(&logger.Config{
	// 	Level:    logger.WarnLevel,
	// 	LogFile:  path + ".warn.log",
	// 	Caller:   true,
	// 	UseColor: true,
	// 	UseJSON:  true,
	// })
}

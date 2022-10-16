package gologutils

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var RootLogger Logging

type Logging struct {
	logger     *zerolog.Logger
	fileLogger *lumberjack.Logger
	debug      *bool
	console    *bool
}

func InitLog(debug bool, console bool, logFileConfig *LogFileConfig) {
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	var targets []io.Writer
	var logFile *lumberjack.Logger
	if logFileConfig != nil {
		logFile = newRollingFile(logFileConfig)
		targets = append(targets, logFile)
	}
	if console || len(targets) == 0 {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		targets = append(targets, consoleWriter)
	}
	multi := zerolog.MultiLevelWriter(targets...)
	log := zerolog.New(multi).With().Timestamp().Logger()
	RootLogger = Logging{
		&log,
		logFile,
		&debug,
		&console,
	}
}

// Close ...
func (l *Logging) Close() {
	if l.fileLogger != nil {
		l.fileLogger.Close()
	}
}

func (l *Logging) NewLogger(componentName string) *zerolog.Logger {
	var ret zerolog.Logger
	if len(componentName) > 0 {
		ret = l.logger.With().Str("component", componentName).Logger()
	} else {
		ret = *l.logger
	}
	return &ret
}

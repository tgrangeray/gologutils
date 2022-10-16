package gologutils

import "gopkg.in/natefinch/lumberjack.v2"

type LogFileConfig struct {
	Filename     string
	MaxSizeMB    int
	MaxBackups   int
	MaxAgeInDays int
}

func (l *LogFileConfig) fillDefaults() {
	if l.MaxAgeInDays == 0 {
		l.MaxAgeInDays = 30
	}
	if l.MaxSizeMB == 0 {
		l.MaxSizeMB = 100
	}
	if l.MaxBackups == 0 {
		l.MaxBackups = 1
	}
}

func newRollingFile(config *LogFileConfig) *lumberjack.Logger {
	l := lumberjack.Logger{
		Filename:   config.Filename,
		MaxBackups: config.MaxBackups,   // files
		MaxSize:    config.MaxSizeMB,    // megabytes
		MaxAge:     config.MaxAgeInDays, // days
	}
	return &l
}

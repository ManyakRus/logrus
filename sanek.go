package logrus

import (
	gorm_logger "gorm.io/gorm/logger"
)

// LogLevel log level
//type LogLevel int

// LogMode log mode
func (l *Logger) LogMode(level gorm_logger.LogLevel) *Logger {

	i := int(level)
	u := uint32(i)
	level1 := Level(u)
	l.SetLevel(level1)

	return l
}

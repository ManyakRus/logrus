package logrus

// LogLevel log level
type LogLevel int

// LogMode log mode
func (l *Logger) LogMode(level LogLevel) *Logger {

	i := int(level)
	u := uint32(i)
	level1 := Level(u)
	l.SetLevel(level1)

	return l
}

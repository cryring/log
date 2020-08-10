package log

type Logger struct {
	appenders []Appender
	maxLevel  Level
}

func NewLogger() *Logger {
	return nil
}

func (l *Logger) AddAppender(appender Appender) {
	l.appenders = append(l.appenders, appender)
}

func (l *Logger) MaxLevel() Level {
	return l.maxLevel
}

func (l *Logger) SetMaxLevel(level Level) {

}

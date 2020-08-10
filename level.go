package log

type Level int8

const (
	None Level = iota
	Fatal
	Error
	Warning
	Info
	Debug
	Verbose
)

func (l *Level) String() string {
	switch *l {
	case Fatal:
		return "FATAL"
	case Error:
		return "ERROR"
	case Warning:
		return "WARN"
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	case Verbose:
		return "VERB"
	default:
		return "NONE"
	}
}

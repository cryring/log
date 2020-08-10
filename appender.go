package log

type Appender interface {
	Write(r *Record) error
}

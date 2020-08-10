package log

import "time"

type Record struct {
	level   Level
	tm      time.Time
	message string
}

func NewRecord() *Record {
	return nil
}

func (r *Record) Level() Level {
	return r.level
}

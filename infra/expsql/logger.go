package expsql

import "log"

// Logger is an interface that reports an extra error.
type Logger interface {
	SetRollbackError(error)
}

type defaultLogger struct{}

func (l *defaultLogger) SetRollbackError(err error) {
	log.Printf("SetRollbackError is `%v`.", err)
}

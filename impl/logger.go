package impl

import (
	"log"
	"os"
)

type BaseLogger struct {
	*log.Logger
}

func errLogger() *BaseLogger {
	errorLog := &BaseLogger{Logger: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)}

	return errorLog
}

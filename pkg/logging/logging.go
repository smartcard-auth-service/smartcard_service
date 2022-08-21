package logging

import (
	"os"
	"sync"

	"log"
)

type Logger struct {
	filename string
	*log.Logger
}

var logger *Logger
var once sync.Once

func GetLogger() (*Logger, error) {
	var errCreateLogger error
	once.Do(func() {
		logger, errCreateLogger = createLogger("logs/service.log")
	})
	return logger, errCreateLogger
}

func createLogger(fname string) (*Logger, error) {
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &Logger{
		filename: fname,
		Logger:   log.New(file, "My app Name ", log.Lshortfile),
	}, err
}

package logging

import (
	"os"
	"sync"

	"log"
)

type CustomLogger struct {
	filename string
	Jrn      *log.Logger
}

var logger *CustomLogger
var once sync.Once

func GetLogger() (*CustomLogger, error) {
	var errCreateLogger error
	once.Do(func() {
		logger, errCreateLogger = createLogger("logs/service.log")
	})
	return logger, errCreateLogger
}

func createLogger(fname string) (*CustomLogger, error) {
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &CustomLogger{
		filename: fname,
		Jrn:      log.New(file, "My app Name ", log.Lshortfile),
	}, err
}
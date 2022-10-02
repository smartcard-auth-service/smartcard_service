package logging

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var Logrus = logrus.New()

func InitLogger() error {
	errCreateLogger := createLogger("logs/service.log")
	return errCreateLogger
}

func createLogger(fname string) error {
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err == nil {
		Logrus.Out = file
	} else {
		Logrus.Info("Failed to Logrus to file, using default stderr")
		return err
	}
	Logrus.SetLevel(logrus.TraceLevel)
	Logrus.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
		},
		FullTimestamp: true,
	})
	Logrus.SetReportCaller(true)
	return nil
}

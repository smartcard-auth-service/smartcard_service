package app

import (
	"smartcard/pkg/logging"
)

type App struct {
	Logger *logging.CustomLogger
}

func NewApp() (App, error) {
	return App{}, nil
}

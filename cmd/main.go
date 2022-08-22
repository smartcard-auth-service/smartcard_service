package main

import (
	"smartcard/internal/app/server"
	log "smartcard/pkg/logging"
)

func main() {
	logger, _ := log.GetLogger()
	logger.Jrn.Println("Starting App")
	server.Run()
}

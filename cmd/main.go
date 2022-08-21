package main

import (
	log "smartcard/pkg/logging"
)

func main() {
	logger, _ := log.GetLogger()
	logger.Println("Starting App")
}

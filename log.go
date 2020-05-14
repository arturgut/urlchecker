package main

import (
	log "github.com/sirupsen/logrus"
)

func initLogging() {

	log.SetLevel(log.DebugLevel) // Only log the warning severity or above.

}

package main

import (
	"github.com/AlexsJones/vinculum/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.DebugLevel)
	cmd.Execute()
}

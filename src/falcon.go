package main

import (
	"client"
	"daemon"
	"flag"
	"fmt"
	log "github.com/zdannar/flogger"
)

const (
	LOGFILE = "/var/log/falcon.log"
)

func setupLogging() {
	if err := log.OpenFile(LOGFILE, log.FLOG_APPEND, 0644); err != nil {
		log.Fatalf("Unable to open log file : %s", err)
	}

	log.SetLevel(log.INFO)
	log.RedirectStreams()
}

func main() {
	flag.String("mode", "client", "falcon mode to run in")

	flag.Parse()
	if flag.NArg() != 1 {
		log.Infof("not sufficient arguments supplied.")
		flag.PrintDefaults()
	}

	setupLogging()

	mode := flag.Arg(0)
	if mode == "daemon" {
		daemon.Daemon()
	}
	if mode == "client" {
		client.Client()
	}
}

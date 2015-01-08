package main

import (
	"client"
	"daemon"
	"flag"
	"log"
)

func main() {
	flag.String("mode", "client", "falcon mode to run in")

	flag.Parse()
	if flag.NArg() != 1 {
		log.Println("not sufficient arguments supplied.")
		flag.PrintDefaults()
	}

	mode := flag.Arg(0)
	if mode == "daemon" {
		port := flag.Arg(1)
		daemon.Daemon(port)
	}

	if mode == "client" {
		hostPort := flag.Arg(1)
		client.Client(hostPort)
	}
}

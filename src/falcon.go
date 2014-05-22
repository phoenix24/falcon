package main

import (
	"fmt"
	"flag"
	"client"
	"daemon"
)

func main() {
	flag.String("mode", "client", "falcon mode to run in")

	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("not sufficient arguments supplied.")
		flag.PrintDefaults()
	}

	mode := flag.Arg(0)
	if mode == "daemon" {
		daemon.Daemon()
	}
	if mode == "client" {
		client.Client()
	}
}

package client

import (
	"fmt"
	log "github.com/zdannar/flogger"
	"net"
	"os"
)

var (
	port        = ":4224"
	application = "falcon-client"
)

func Client() {
	log.Infof("%s, running., connecting to localhost on port %s\n", application, port)

	address, err := net.ResolveTCPAddr("tcp", port)
	handleError(err)

	conn, err := net.DialTCP("tcp", nil, address)
	handleError(err)

	message := "hello world"
	request := []byte(fmt.Sprintf("request: %s\n", message))
	conn.Write(request)

	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		handleError(err)

		response := buf[0:n]
		log.Infof(string(response))

		conn.Write(response[0:])
	}
}

func handleError(err error) {
	if err != nil {
		log.Infof("error while bootstrapping. abort. %s\n", err)
		os.Exit(1)
	}
}

package client

import (
	"fmt"
	"log"
	"net"
	"os"
)

func Client(hostPort string) {
	log.Println("falcon-client connecting to %s\n", hostPort)

	address, err := net.ResolveTCPAddr("tcp", hostPort)
	handleError(err)

	conn, err := net.DialTCP("tcp", nil, address)
	handleError(err)

	message := "hello world"
	request := []byte(fmt.Sprintf("request: %s\n", message))
	conn.Write(request)

	var buf [512]byte

	n, err := conn.Read(buf[0:])
	handleError(err)

	response := buf[0:n]
	log.Println(string(response))

	conn.Write(response[0:])
}

func handleError(err error) {
	if err != nil {
		log.Println("error while bootstrapping. abort. %s\n", err)
		os.Exit(1)
	}
}

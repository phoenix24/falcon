package daemon

import (
	log "github.com/zdannar/flogger"
	"net"
	"os"
)

var (
	port        = ":4224"
	application = "falcon-daemon"
)

func Daemon() {
	log.Infof("%s, listening on port %s\n", application, port)

	address, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		handleError(err)
	}

	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		handleError(err)
	}

	for {
		connection, err := listener.AcceptTCP()
		if err != nil {
			continue
		}

		go handleClient(connection)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		log.Infof(string(buf[0:n]))
		_, err1 := conn.Write(buf[0:n])
		if err1 != nil {
			return
		}
	}
}

func handleError(err error) {
	log.Infof("Error while bootstrapping. Aborting. %s\n", err)
	os.Exit(1)
}

package daemon

import (
	"fmt"
	"net"
	"os"
)

var (
	port = ":4224"
	application = "falcon-daemon"
)

func Daemon() {
	fmt.Printf("%s, listening on port %s\n", application, port)

	address, err := net.ResolveTCPAddr("tcp", port)
	handleError(err)

	listener, err := net.ListenTCP("tcp", address)
	handleError(err)

	for {
		var connection, err = listener.Accept()
		if err != nil { continue }

		handleClient(connection)
	}
}

func handleClient(conn net.Conn) {
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil { return }

		fmt.Printf(string(buf[0:n]))
		_, err1 := conn.Write(buf[0:n])
		if err1 != nil { return }
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("error while bootstrapping. abort. %s\n", err)
		os.Exit(1)
	}
}

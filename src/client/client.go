package client

import "fmt"


var (
	port = 4224
	application = "falcon-daemon"
)

func Client() {
	fmt.Printf("%s, running on port: %d", application, port)
}

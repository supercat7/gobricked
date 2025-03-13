package comms

import (
	"fmt"
	"net"
)

func authClient(cli net.Conn, pass string) {
	var input = make([]byte, 1024)
	n, err := cli.Read(input)
	if err != nil {
		fmt.Println("Failed to read from operator:", err)
	}

	var val string = string(input[:n])

	if val == pass {
		fmt.Println("Operator authenticated-->", cli.RemoteAddr())
		cli.Write([]byte(string("OK")))
	} else {
		fmt.Println("Operator could not be authenticated, bad password:", pass)
		cli.Write([]byte(string("ERR")))
	}
}

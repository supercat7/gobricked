package comms

import (
	"fmt"
	"net"
)

func authClient(cli net.Conn) {
	var input = make([]byte, 1024)
	n, err := cli.Read(input)
	if err != nil {
		fmt.Println("Failed to authenticate client:", err)
	}
	fmt.Printf("Here is the input from client: %s\n", string(input[:n]))
}

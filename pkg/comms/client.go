package comms

import (
	"fmt"
	"net"
)

func authClient(cli net.Conn) {
	var input []byte
	_, err := cli.Read(input)
	if err != nil {
		fmt.Println("Failed to authenticate client:", cli)
	}
	fmt.Printf("Here is the input from client: %s\n", string(input))
}

package comms

import (
	"fmt"
	"gobricked/client/comms"
	"net"
	"strings"
)

func authClient(cli net.Conn, userPass map[string]string) {
	var input = make([]byte, 1024)
	n, err := cli.Read(input)
	if err != nil {
		fmt.Println("Failed to authenticate operator:", err)
	}
	msg := strings.Split(string(input[:n]), ":")
	user, pass := msg[0], msg[1]
	val, ok := userPass[user]
	if ok {
		if val == pass {
			fmt.Println("Operator authenticated-->", user+":"+pass)
			comms.AUTHENTICATED = true
			cli.Write([]byte(string("OK")))
		} else {
			fmt.Println("Operator could not be authenticated, bad password:", user+":"+pass)
			cli.Write([]byte(string("ERR")))
		}
	} else {
		fmt.Println("Operator could not be authenticated, bad username or password:", user+":"+pass)
		cli.Write([]byte(string("ERR")))
	}
}

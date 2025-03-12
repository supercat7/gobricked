package comms

import (
	"fmt"
	"net"
)

var AUTHENTICATED bool = false

type SockStream struct {
	IP   string
	Port string
	Conn net.Conn
}

func (s *SockStream) AuthServer(username string, password string) {
	var err error
	s.Conn, err = net.Dial("tcp", s.IP+":"+s.Port)
	if err != nil {
		panic(err)
	}
	fmt.Println("Established connection to:", s.Conn.RemoteAddr())
	var msg string = username + ":" + password
	fmt.Println("Sending username and password...")
	s.Conn.Write([]byte(msg))

	var recv = make([]byte, 1024)
	n, e := s.Conn.Read(recv)
	if e != nil {
		fmt.Println("Error receiving authentication confirmation from server:", e)
		return
	}
	if string(recv[:n]) == "OK" {
		fmt.Println("OK code received, successfully authenticated with server")
		AUTHENTICATED = true
	} else if string(recv[:n]) == "ERR" {
		fmt.Println("ERR code received, bad username or password, authentication failure..")
	}
}

package comms

import (
	"fmt"
	"net"
	"strings"
)

type SockStream struct {
	IP   string
	Port string
	Conn net.Conn
}

func NewSockStream(addr string) *SockStream {
	arg := strings.Split(addr, ":")
	var s *SockStream = &SockStream{
		IP:   arg[0],
		Port: arg[1],
	}
	return s
}

func (s *SockStream) AuthServer(password string) bool {
	var err error
	s.Conn, err = net.Dial("tcp", s.IP+":"+s.Port)
	if err != nil {
		panic(err)
	}
	fmt.Println("Established connection to:", s.Conn.RemoteAddr())
	fmt.Println("Sending password...")
	s.Conn.Write([]byte(password))

	var recv = make([]byte, 1024)
	n, e := s.Conn.Read(recv)
	if e != nil {
		fmt.Println("Error receiving authentication confirmation from server:", e)
		return false
	}
	if string(recv[:n]) == "OK" {
		fmt.Println("OK code received, successfully authenticated with server")
		return true
	} else if string(recv[:n]) == "ERR" {
		fmt.Println("ERR code received, bad username or password, authentication failure..")
		return false
	}
	return false
}

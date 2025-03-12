package comms

import "net"

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
	var msg string = username + ":" + password
	s.Conn.Write([]byte(msg))
}

package comms

import (
	"fmt"
	"net"
	"sync"
)

type TCPServer struct {
	Port     string
	Listener net.Listener
	Connec   net.Conn
	Running  bool
	mu       sync.Mutex
}

func NewTCPServer(port string) *TCPServer {
	return &TCPServer{
		Port: port,
	}
}

func (t *TCPServer) Start() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.Running {
		fmt.Printf("Server already running on port: %s\n", t.Port)
	}

	var err error
	t.Listener, err = net.Listen("tcp", ":"+t.Port)
	if err != nil {
		fmt.Printf("\nErr: Failed to bind server to port: %v\n", err)
		return
	}
	t.Running = true
	fmt.Println("\nServer started on port:", t.Port)
	for {
		t.Connec, err = t.Listener.Accept()
		if err != nil {
			fmt.Printf("\nErr: Failed to accept client connection: %v\n", err)
		}
		fmt.Printf("\nConnection received from: %s\n", t.Connec.RemoteAddr())
		go HandleClientConnection(t.Connec)
	}
}

func (t *TCPServer) Stop() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.Running {
		fmt.Printf("\nServer is not running on port: %s\n", t.Port)
		return
	}

	if t.Listener != nil {
		t.Listener.Close()
		t.Running = false
		fmt.Println("\nClosed server on port:", t.Port)
	} else {
		fmt.Println("\nServer is not running on port:", t.Port)
	}
}

func StopAllServers() {}

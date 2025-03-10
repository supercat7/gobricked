package comms

import (
	"fmt"
	"gobricked/pkg/server/art"
	"net"
	"sync"
)

var AGENT_SOCK_LIST []net.Conn

type TCPServer struct {
	Port     string
	Listener net.Listener
	Connec   net.Conn
	Running  bool
	mu       sync.Mutex
}

var SERVERINSTANCE *TCPServer = NewTCPServer("9090")
var SERVERCHANNEL chan struct{} = make(chan struct{})

func NewTCPServer(port string) *TCPServer {
	return &TCPServer{
		Port: port,
	}
}

func (t *TCPServer) Start(quit chan struct{}) {
	t.mu.Lock()

	if t.Running {
		fmt.Printf("Server already running on port: %s\n", t.Port)
		t.mu.Unlock()
		return
	}

	var err error
	t.Listener, err = net.Listen("tcp", ":"+t.Port)
	if err != nil {
		fmt.Printf("\nErr: Failed to bind server to port: %v\n", err)
		t.mu.Unlock()
		return
	}
	t.Running = true

	t.mu.Unlock()
	for {
		select {
		case <-quit:
			fmt.Println("Shutting Listener...")
			t.Listener.Close()
			return
		default:
			t.Connec, err = t.Listener.Accept()
			if err != nil {
				fmt.Printf("\nErr: Failed to accept client connection: %v\n", err)
				fmt.Print(art.GOBRICKED_PROMPT)
			}
			fmt.Printf("\nConnection received from: %s\n\r", t.Connec.RemoteAddr())
			fmt.Print(art.GOBRICKED_PROMPT)
			AGENT_SOCK_LIST = append(AGENT_SOCK_LIST, t.Connec)
		}
	}
}

func StopAllAgentComms() {
	for i := 0; i < len(AGENT_SOCK_LIST); i++ {
		fmt.Println("Closed connection to:", AGENT_SOCK_LIST[i].RemoteAddr())
		AGENT_SOCK_LIST[i].Close()
	}
}

func (t *TCPServer) Stop(quit chan struct{}) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.Running {
		fmt.Println("\nServer is not running on port:", t.Port)
		t.mu.Unlock()
		return
	}
	close(quit)

	if t.Listener != nil {
		t.Running = false
		fmt.Println("Sending shutdown signal to all agents...")
		StopAllAgentComms()
		fmt.Println("\nClosed server on port:", t.Port)
	} else {
		fmt.Println("\nServer is not running on port:", t.Port)
	}
}

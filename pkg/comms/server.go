package comms

import (
	"fmt"
	"gobricked/pkg/util"
	"net"
	"sync"
)

type OperatorServer struct {
	Port    string
	Server  net.Listener
	Connec  net.Conn
	Running bool
	mu      sync.Mutex
	config  util.ServerConfig
}

type Server interface {
	Start()
	Stop()
}

func NewOperatorServer(port string, c util.ServerConfig) *OperatorServer {
	return &OperatorServer{
		Port:   port,
		config: c,
	}
}

func (t *OperatorServer) Start(quit chan struct{}) {
	t.mu.Lock()

	if t.Running {
		fmt.Printf("Server already running on port: %s\n", t.Port)
		t.mu.Unlock()
		return
	}

	var err error
	t.Server, err = net.Listen("tcp", ":"+t.Port)
	if err != nil {
		fmt.Printf("\nErr: Failed to bind server to port: %v\n", err)
		t.mu.Unlock()
		return
	}
	t.Running = true
	fmt.Println("Waiting for connections from operators...")
	t.mu.Unlock()
	for {
		select {
		case <-quit:
			fmt.Println("Shutting OperatorServer...")
			t.Server.Close()
			return
		default:
			t.Connec, err = t.Server.Accept()
			if err != nil {
				fmt.Printf("\nErr: Failed to accept operator connection: %v\n", err)
			}
			fmt.Printf("\nConnection received from: %s\n\r", t.Connec.RemoteAddr())
			fmt.Println("Attempting to authenticate operator...")
			go authClient(t.Connec, t.config.Teamserver.Password)
		}
	}
}

func (t *OperatorServer) Stop(quit chan struct{}) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.Running {
		fmt.Println("\nServer is not running on port:", t.Port)
		t.mu.Unlock()
		return
	}
	close(quit)

	if t.Server != nil {
		t.Running = false
		fmt.Println("Closed server on port:", t.Port)
	} else {
		fmt.Println("Server is not running on port:", t.Port)
	}
}

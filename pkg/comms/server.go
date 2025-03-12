package comms

import (
	"fmt"
	"gobricked/pkg/util"
	"net"
	"sync"
)

type Listener struct {
	Port     string
	Listener net.Listener
	Connec   net.Conn
	Running  bool
	mu       sync.Mutex
}

type Server interface {
	Start()
	Stop()
}

func NewListener(port string) *Listener {
	return &Listener{
		Port: port,
	}
}

func (t *Listener) Start(quit chan struct{}, config util.ServerConfig) {
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
	fmt.Println("Waiting for connections from operators...")
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
				fmt.Printf("\nErr: Failed to accept operator connection: %v\n", err)
			}
			fmt.Printf("\nConnection received from: %s\n\r", t.Connec.RemoteAddr())
			fmt.Println("Attempting to authenticate operator...")
			authClient(t.Connec, util.GetOperators(config))
		}
	}
}

func (t *Listener) Stop(quit chan struct{}) {
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
		fmt.Println("Closed server on port:", t.Port)
	} else {
		fmt.Println("Server is not running on port:", t.Port)
	}
}

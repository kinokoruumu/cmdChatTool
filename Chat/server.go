package Chat

import (
	"net"
	"fmt"
)

func Server(port string) {
	address := "0.0.0.0:" + port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Listen error: %s\n", err)
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("Accept error: %s\n", err)
		return
	}
	defer conn.Close()

	go Listen(conn)
	Input(conn)
}

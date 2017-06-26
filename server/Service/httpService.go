package Service

import (
	"../Constants"
	"net"
	"fmt"
)

func HttpService() {
	address := Constants.SERVER_IPADDR + Constants.SERVER_PORT
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Listen error: %s\n", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept error: %s\n", err)
			return
		}
		//defer conn.Close()

		go ReadHandle(conn)
		go SendHandle(conn)
	}

	//if conn != nil {
	//	listener, err := net.Listen("tcp", address)
	//	if err != nil {
	//		fmt.Printf("Listen error: %s\n", err)
	//		return
	//	}
	//	defer listener.Close()
	//}

	//var conn []net.Conn
	//conn = append(conn, listener.Accept())
	//conn, err := listener.Accept()
	//if err != nil {
	//	fmt.Printf("Accept error: %s\n", err)
	//	return
	//}
	//defer conn.Close()
	//
	//go ReadHandle(conn)
	//SendHandle(conn)
}

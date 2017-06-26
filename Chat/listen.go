package Chat

import (
	"net"
	"fmt"
)

func Listen(conn net.Conn)  {
	fmt.Println("メッセージ:")
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Read error: %s\n", err)
		}
		if n == 0 {
			break
		}
		fmt.Print("[受信] " + string(buf[:n])+"\n")
	}
}
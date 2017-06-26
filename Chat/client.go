package Chat

import (
	"fmt"
	"net"
)

func Client(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("サーバーが見つかりませんでした\n")
		return
	}
	defer conn.Close()

	go Listen(conn)
	Input(conn)
}
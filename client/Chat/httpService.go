package Chat

import (
	"../Constants"
	"fmt"
	"net"
)

func HttpService() {
	address := Constants.CLIENT_IPADDR + Constants.CLIENT_PORT
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("サーバーが見つかりませんでした\n")
		return
	}
	//defer conn.Close()

	go ReadHandle(conn)
	SendHandle(conn)
}
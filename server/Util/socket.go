package Util

import (
	"net"
	"fmt"
)

type Socket struct {
	conn net.Conn
}

func SocketManager(conn net.Conn) *Socket{
	return &Socket{conn: conn}
}

func (this *Socket)Send(msg string){
	this.conn.Write([]byte(msg))
}

func (this *Socket)Read(buf []byte) string{
	n, err := this.conn.Read(buf)
	if err != nil {
		fmt.Printf("Read error: %s\n", err)
	}
	return string(buf[:n])
}

func (this *Socket)Close(){
	this.conn.Close()
}

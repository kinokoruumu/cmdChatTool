package Chat

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func ReadHandle(conn net.Conn)  {
	fmt.Println("メッセージ:")
	buf := make([]byte, 1024)
	sock := SocketManager(conn)
	for {
		n := sock.Read(buf)
		if n == "" {
			break
		}
		fmt.Print("[受信]: " + n +"\n")
	}
}

func SendHandle(conn net.Conn)  {
	s := bufio.NewScanner(os.Stdin)
	sock := SocketManager(conn)
	for s.Scan() {
		input := s.Text()
		if input == "exit"{
			sock.Close()
			break
		}
		sock.Send(input)
		fmt.Print("[送信]: " + input +"\n")
	}
	defer sock.Close()
}
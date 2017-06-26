package Chat

import (
	"bufio"
	"os"
	"net"
)

func Input(conn net.Conn) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		input := s.Text()
		conn.Write([]byte(input))
	}
}
package main

import (
	"./server/Service"
	"./client/Chat"
	"os"
	"fmt"
	"flag"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
		Usage of %s:
		   %s [OPTIONS] ARGS...
		Options\n`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	var (
		server = flag.Bool("s", false, "up to server")
		client = flag.Bool("c", false, "up to client")
	)
	flag.Parse()

	if *server {
		Service.HttpService()
	}

	if *client {
		Chat.HttpService()
	}
}
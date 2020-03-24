package main

import (
	"fmt"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":40080")
	if err != nil {
		fmt.Println("Error while Listening -> ", err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("Error while Accepting connection -> ", err)
			continue
		}
		go handle(conn)
	}
}

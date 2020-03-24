package main

import (
	"encoding/csv"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	var record []string
	var err error

	//csv reads and writes comma-separated values
	cv := csv.NewReader(conn)
	
}
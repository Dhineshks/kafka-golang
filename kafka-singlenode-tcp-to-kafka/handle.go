package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	var record []string
	var err error

	//csv reads and writes comma-separated values
	cv := csv.NewReader(conn)

	//recover is builtin function which is used to regain the control of panicking goroutine
	//recover is useful only when called inside defered function
	//executing a call to recover inside a defered function stops the paniking sequence by restoring
	//normal execution and retrives the error value passed to the call of panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered for -> ", r)
		}
	}()
	for {
		if record, err = cv.Read(); err != nil || len(record) != 5 {
			if err == io.EOF || len(record) == 0 {
				fmt.Println("Empty record")
				return
			}
			return
		}
		fmt.Println(record)
	}

}

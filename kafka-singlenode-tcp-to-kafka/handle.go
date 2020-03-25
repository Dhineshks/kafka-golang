package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/segmentio/kafka-go"
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
		if record, err = cv.Read(); err != nil || len(record) != 6 {
			if err == io.EOF || len(record) == 0 {
				fmt.Println("Empty record")
				return
			}
			return
		}

		name := record[0]
		loc := record[1]
		ph := record[2]
		id := record[3]
		ti := record[4]
		amount := record[5]
		d := time.Now()
		t := d.Format("2006.01.02")

		st := Info{
			CustomerName: name,
			Locality:     loc,
			Phone:        ph,
			CustomerID:   id,
			Date:         t,
			Data: []Datastruct{
				Datastruct{
					Timestamp:   ti,
					TotalAmount: amount,
				},
			},
		}
		kafkastream(st)
	}

}

func kafkastream(x Info) {
	//encoding json
	js, err := json.Marshal(x)
	if err != nil {
		fmt.Println("can't Decode", err)
	}
	jsstring := string(js)

	topic := "customer-data"
	partition := 0

	con, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	con.SetWriteDeadline(time.Now().Add(20 * time.Second))

	for _, word := range []string{string(jsstring)} {
		con.WriteMessages(
			kafka.Message{Value: []byte(word)},
		)
	}
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "customer-data"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	defer conn.Close()

	//setting 10kb minimum and 1mb maximum
	batch := conn.ReadBatch(10e3, 1e6)
	defer batch.Close()
	//10kb per message
	b := make([]byte, 10e3)
	for {
		_, err := batch.Read(b)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
	}
}

func mongocli(m []byte) {
	var parse Info
	json.Unmarshal([]byte(m), &parse)
}

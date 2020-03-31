package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "customer-data"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

}

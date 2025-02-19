package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbCreds = options.Client().ApplyURI("mongodb://localhost:27017")

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

func getclient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client(), dbCreds)
	if err != nil {
		log.Fatal(err)
	}
	if err == nil {
		return client, err
	}
	return client, err
}

func mongocli(m []byte) {
	var a Info
	json.Unmarshal([]byte(m), &a)

	cli, err := getclient()
	if err != nil {
		fmt.Println(err)
	}
	defer cli.Disconnect(context.Background())
	collection := cli.Database("Data").Collection("customer")
}

package kafkalib

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func InitKafka() *kafka.Conn {
	topic := "test_topic"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	return conn
}

func WriteMessageOnTopic(conn *kafka.Conn) {
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}

func ReadMessageFromTopic(conn *kafka.Conn) {

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e0, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}
}

func CloseKafkaConnection(conn *kafka.Conn) {
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

package main

import (
	"fmt"
	kafkalib "learn-kafka/kafka-lib"
)

func main() {
	fmt.Println("Hello kafka")

	conn := kafkalib.InitKafka()
	kafkalib.WriteMessageOnTopic(conn)
	fmt.Println("Reading ......")
	kafkalib.ReadMessageFromTopic(conn)
	kafkalib.CloseKafkaConnection(conn)

	fmt.Println("Executed")
}

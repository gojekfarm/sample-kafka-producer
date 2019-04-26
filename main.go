package main

import (
	"flag"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gogo/protobuf/proto"
	msg "github.com/gojekfarm/sample-kafka-producer/proto"
)

func main() {
	limit := flag.Int("limit", 10, "total messages")
	start := flag.Int("start", 0, "start of order number")
	//TODO: pull out host, flush wait, and topic as flags

	flag.Parse()

	config := &kafka.ConfigMap{"bootstrap.servers": "localhost:9092"}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	for i := *start; i < *limit; i++ {
		orderNo := fmt.Sprintf("order-%d", i)
		message := msg.TestMessage{
			OrderNumber: orderNo,
			OrderUrl:    fmt.Sprintf("Url-%d", i),
			Success:     true,
		}
		topic := "test-messages-01"

		value, err := proto.Marshal(&message)
		if err != nil {
			panic(err)
		}
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          value,
		}, nil)
		fmt.Printf("sent message: %s\n", orderNo)
	}
	fmt.Println("waiting for producer to flush data...")
	producer.Flush(2 * 1000)
}

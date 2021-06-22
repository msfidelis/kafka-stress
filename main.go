package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	guuid "github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

func main() {
	topic := flag.String("topic", "kafka-stress", "Kafka Stress Topics")
	createTopic := flag.String("create-topic", "false", "Auto Create Topic?")
	testMode := flag.String("test-mode", "producer", "Test Type; Ex producer;consumer. Default: producer")
	bootstrapServers := flag.String("bootstrap-servers", "0.0.0.0:9092", "Kafka Bootstrap Servers Broker Lists")
	zookeeperServers := flag.String("zookeeper-servers", "0.0.0.0:2181", "Zookeeper Connection String")
	events := flag.Int("events", 10000, "Numer of events will be created in topic")

	flag.Parse()

	fmt.Println(*events)
	fmt.Println(*topic)
	fmt.Println(*testMode)
	fmt.Println(*createTopic)
	fmt.Println(*bootstrapServers)
	fmt.Println(*zookeeperServers)

	switch strings.ToLower(*testMode) {
	case "producer":
		produce(*bootstrapServers, *topic, *events)
		break
	case "consumer":
		consume()
		break
	default:
		return
	}
}

func produce(bootstrap_servers string, topic string, events int) {
	producer := getProducer(bootstrap_servers, topic)

	defer producer.Close()

	var wg sync.WaitGroup

	var executions uint64
	var errors uint64

	for i := 0; i < events; i++ {
		wg.Add(1)

		go func() {

			msg := kafka.Message{
				Key:   []byte(guuid.New().String()),
				Value: []byte(guuid.New().String()),
			}

			err := producer.WriteMessages(context.Background(), msg)

			if err != nil {
				fmt.Println(err)
				atomic.AddUint64(&errors, 1)
			} else {
				atomic.AddUint64(&executions, 1)
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Sent %v messages to topic %s with %v errors \n", executions, topic, errors)
}

func consume() {

}

func getProducer(bootstrap_servers, topic string) *kafka.Writer {

	name, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	dialer := &kafka.Dialer{
		Timeout:   20 * time.Second,
		DualStack: true,
		ClientID:  name,
	}

	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: strings.Split(bootstrap_servers, ","),
		Topic:   topic,
		// Balancer: &kafka.LeastBytes{},
		Balancer:     &kafka.Hash{},
		Dialer:       dialer,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	})

}

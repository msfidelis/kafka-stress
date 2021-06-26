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
	createTopic := flag.Bool("create-topic", false, "Auto Create Topic?")
	testMode := flag.String("test-mode", "producer", "Test Type; Ex producer;consumer. Default: producer")
	bootstrapServers := flag.String("bootstrap-servers", "0.0.0.0:9092", "Kafka Bootstrap Servers Broker Lists")
	zookeeperServers := flag.String("zookeeper-servers", "0.0.0.0:2181", "Zookeeper Connection String")
	events := flag.Int("events", 10000, "Numer of events will be created in topic")
	consumers := flag.Int("consumers", 1, "Number of consumers will be used in topic")
	consumerGroup := flag.String("consumer-group", "kafka-stress", "Consumer group name")

	flag.Parse()

	if *createTopic {
		createTopicAfterTest(*topic, *zookeeperServers)
	}

	switch strings.ToLower(*testMode) {
	case "producer":
		produce(*bootstrapServers, *topic, *events)
		break
	case "consumer":
		consume(*bootstrapServers, *topic, *consumerGroup, *consumers)
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

	start := time.Now()

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
	elapsed := time.Since(start)
	meanEventsSent := float64(executions) / elapsed.Seconds()
	fmt.Printf("Sent %v messages to topic %s with %v errors \n", executions, topic, errors)
	fmt.Printf("Tests finished in %v. Producer mean time %.2f/s \n", elapsed, meanEventsSent)
}

func consume(bootstrap_servers, topic, consumer_group string, consumers int) {

	var wg sync.WaitGroup

	for i := 0; i < consumers; i++ {
		fmt.Println(i)
		consumerId := i + 1
		consumer := getConsumer(bootstrap_servers, topic, consumer_group, consumerId)
		wg.Add(1)
		go func() {
			for {
				m, err := consumer.ReadMessage(context.Background())
				if err != nil {
					break
				}
				fmt.Printf("[Client %v] message from consumer group %s at topic/partition/offset %v/%v/%v: %s = %s\n", consumerId, consumer_group, m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
			}
			wg.Done()
		}()
	}

	// consumer := getConsumer(bootstrap_servers, topic, consumer_group)

	// for {
	// 	m, err := consumer.ReadMessage(context.Background())
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Printf("message from consumer group %s at topic/partition/offset %v/%v/%v: %s = %s\n", consumer_group, m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	// }

	fmt.Println("Consumer")
}

func createTopicAfterTest(topic string, zookeeper string) {
	fmt.Printf("Creating topic %s\n", topic)
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

func getConsumer(bootstrap_servers, topic, consumer_group string, consumer int) *kafka.Reader {

	// @TODO Separar um dialer
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		ClientID:  fmt.Sprintf("%v-%v", consumer_group, consumer),
	}

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: strings.Split(bootstrap_servers, ","),
		Topic:   topic,
		GroupID: consumer_group,
		Dialer:  dialer,
	})

}

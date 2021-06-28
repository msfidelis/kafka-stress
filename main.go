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

	"kafka-stress/pkg/string_generator"

	guuid "github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

func main() {
	topic := flag.String("topic", "kafka-stress", "Kafka Stress Topics")
	createTopic := flag.Bool("create-topic", false, "Auto Create Topic?")
	testMode := flag.String("test-mode", "producer", "Test Type; Ex producer;consumer. Default: producer")
	bootstrapServers := flag.String("bootstrap-servers", "0.0.0.0:9092", "Kafka Bootstrap Servers Broker Lists")
	zookeeperServers := flag.String("zookeeper-servers", "0.0.0.0:2181", "Zookeeper Connection String")
	schemaRegistryURL := flag.String("schema-registry", "0.0.0.0:8081", "Schema Registry URL")
	size := flag.Int("size", 62, "Message size in bytes")
	schema := flag.String("schema", "", "Schema")
	events := flag.Int("events", 10000, "Numer of events will be created in topic")
	consumers := flag.Int("consumers", 1, "Number of consumers will be used in topic")
	consumerGroup := flag.String("consumer-group", "kafka-stress", "Consumer group name")

	flag.Parse()

	if *createTopic {
		createTopicBeforeTest(*topic, *zookeeperServers)
	}

	switch strings.ToLower(*testMode) {
	case "producer":
		produce(*bootstrapServers, *topic, *events, *size, *schemaRegistryURL, *schema)
		break
	case "consumer":

		for i := 0; i < *consumers; i++ {
			var consumerID = i + 1
			go consume(*bootstrapServers, *topic, *consumerGroup, consumerID)
		}

		consume(*bootstrapServers, *topic, *consumerGroup, *consumers)
		break
	default:
		return
	}
}

func produce(bootstrapServers string, topic string, events int, size int, schemaRegistryURL string, schema string) {

	var wg sync.WaitGroup
	var executions uint64
	var errors uint64

	producer := getProducer(bootstrapServers, topic)
	defer producer.Close()

	start := time.Now()

	message := string_generator.RandStringBytes(size)

	for i := 0; i < events; i++ {
		wg.Add(1)

		go func() {

			msg := kafka.Message{
				Key:   []byte(guuid.New().String()),
				Value: []byte(message),
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

func consume(bootstrapServers, topic, consumerGroup string, consumerID int) {
	consumer := getConsumer(bootstrapServers, topic, consumerGroup, consumerID)

	for {
		m, err := consumer.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("[Consumer %v] Message from consumer group %s at topic/partition/offset %v/%v/%v: %s = %s\n", consumerID, consumerGroup, m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	fmt.Println("Consumer", consumerID)
}

func createTopicBeforeTest(topic string, zookeeper string) {
	fmt.Printf("Creating topic %s\n", topic)
}

func getProducer(bootstrapServers, topic string) *kafka.Writer {

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
		Brokers:  strings.Split(bootstrapServers, ","),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		// Balancer:     &kafka.Hash{},
		Dialer:       dialer,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	})

}

func getConsumer(bootstrapServers, topic, consumerGroup string, consumer int) *kafka.Reader {

	// @TODO Separar um dialer
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		ClientID:  fmt.Sprintf("%v-%v", consumerGroup, consumer),
	}

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: strings.Split(bootstrapServers, ","),
		Topic:   topic,
		GroupID: consumerGroup,
		Dialer:  dialer,
	})

}

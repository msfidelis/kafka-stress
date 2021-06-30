package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"kafka-stress/pkg/clients"
	"kafka-stress/pkg/stringgenerator"

	guuid "github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

func main() {
	topic := flag.String("topic", "kafka-stress", "Kafka Stress Topics")
	createTopic := flag.Bool("create-topic", false, "Auto Create Topic?")
	ssl := flag.Bool("ssl-enabled", false, "SSL Mode")
	testMode := flag.String("test-mode", "producer", "Test Type; Ex producer;consumer. Default: producer")
	bootstrapServers := flag.String("bootstrap-servers", "0.0.0.0:9092", "Kafka Bootstrap Servers Broker Lists")
	zookeeperServers := flag.String("zookeeper-servers", "0.0.0.0:2181", "Zookeeper Connection String")
	schemaRegistryURL := flag.String("schema-registry", "0.0.0.0:8081", "Schema Registry URL")
	size := flag.Int("size", 62, "Message size in bytes")
	batchSize := flag.Int("batch-size", 0, "Batch size for producer mode")
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
		produce(*bootstrapServers, *topic, *events, *size, *batchSize, *schemaRegistryURL, *schema, *ssl)
		break
	case "consumer":
		consume(*bootstrapServers, *topic, *consumerGroup, *consumers, *ssl)
	default:
		return
	}
}

func produce(bootstrapServers string, topic string, events int, size int, batchSize int, schemaRegistryURL string, schema string, ssl bool) {

	var wg sync.WaitGroup
	var executions uint64
	var errors uint64

	producer := clients.GetProducer(bootstrapServers, topic, batchSize, ssl)
	defer producer.Close()

	start := time.Now()
	message := stringgenerator.RandStringBytes(size)

	for i := 0; i < events; i++ {
		wg.Add(1)

		go func() {
			msg := kafka.Message{
				Key:   []byte(guuid.New().String()),
				Value: []byte(message),
			}

			err := producer.WriteMessages(context.Background(), msg)
			if err != nil {
				// fmt.Println(err)
				atomic.AddUint64(&errors, 1)
			} else {
				atomic.AddUint64(&executions, 1)
			}

			var multiple = executions % 1000
			if multiple == 0 && executions != 0 {
				fmt.Printf("Sent %v messages to topic %s with %v errors \n", executions, topic, errors)
			}

			wg.Done()
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	meanEventsSent := float64(executions) / elapsed.Seconds()

	fmt.Printf("Tests finished in %v. Produce %v messages with mean time %.2f/s \n", elapsed, executions, meanEventsSent)
}

func consume(bootstrapServers, topic, consumerGroup string, consumers int, ssl bool) {

	var wg sync.WaitGroup
	var counter uint64

	for i := 0; i < consumers; i++ {
		wg.Add(1)
		var consumerID = i + 1
		consumer := clients.GetConsumer(bootstrapServers, topic, consumerGroup, consumerID, ssl)
		consumerName := fmt.Sprintf("%v-%v", consumerGroup, consumerID)

		fmt.Printf("[Consumer %v] Starting consumer\n", consumerName)

		go func() {
			for {
				m, err := consumer.ReadMessage(context.Background())
				if err != nil {
					wg.Done()
					break
				}
				atomic.AddUint64(&counter, 1)
				fmt.Printf("[Consumer %v] Message from consumer group %s at topic/partition/offset %v/%v/%v: %v\n", consumerName, consumerGroup, m.Topic, m.Partition, m.Offset, counter)
			}

			wg.Done()
		}()

	}
	wg.Wait()
}

func createTopicBeforeTest(topic string, zookeeper string) {
	fmt.Printf("Creating topic %s\n", topic)
}

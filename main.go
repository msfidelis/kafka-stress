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
	"kafka-stress/pkg/fakejson"

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
	acks := flag.Int("ack", 1, "Required ACKs to produce messages")
	batchSize := flag.Int("batch-size", 0, "Batch size for producer mode")
	schema := flag.String("schema", "", "Schema")
	events := flag.Int("events", 10000, "Numer of events will be created in topic")
	consumers := flag.Int("consumers", 1, "Number of consumers will be used in topic")
	consumerGroup := flag.String("consumer-group", "kafka-stress", "Consumer group name")
	format :=  flag.String("format", "string", "Events Format; ex string,json,avro")
	verbose := flag.Bool("verbose", false, "Verbose Mode; It Prints Events consumed")
	balancer := flag.String("balancer", "hash", "Balance algorithm for producer mode; Ex: hash,murmur2,crc32")


	flag.Parse()

	if *createTopic {
		createTopicBeforeTest(*topic, *zookeeperServers)
	}

	switch strings.ToLower(*testMode) {
	case "producer":
		produce(*bootstrapServers, *topic, *events, *size, *batchSize, *acks, *schemaRegistryURL, *schema, *ssl, *format, *balancer)
		break
	case "consumer":
		consume(*bootstrapServers, *topic, *consumerGroup, *consumers, *ssl, *verbose)
	default:
		return
	}
}

func produce(bootstrapServers string, topic string, events int, size int, batchSize int, acks int, schemaRegistryURL string, schema string, ssl bool, format string, balancer string) {

	var wg sync.WaitGroup
	var executions uint64
	var errors uint64
	var message string

	producer := clients.GetProducer(bootstrapServers, topic, batchSize, acks, ssl, balancer)
	defer producer.Close()

	start := time.Now()

	for i := 0; i < events; i++ {
		wg.Add(1)

		switch format {
		case "string":
			message = stringgenerator.RandStringBytes(size)
			break;
		case "json":
			message = fakejson.RandJSONPayload()
			break; 
		default:
			message = stringgenerator.RandStringBytes(size)
		}

		go func() {
			msg := kafka.Message{
				Key:   []byte(guuid.New().String()),
				Value: []byte(message),
			}

			err := producer.WriteMessages(context.Background(), msg)
			if err != nil {
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

	fmt.Printf("Tests finished in %v. Produce %v messages with mean time %.2f/s using %s balance algorithm \n", elapsed, executions, meanEventsSent, balancer)
}

func consume(bootstrapServers, topic, consumerGroup string, consumers int, ssl bool, verbose bool) {

	var wg sync.WaitGroup
	var counter uint64

	for i := 0; i < consumers; i++ {
		wg.Add(1)
		var consumerID = i + 1
		consumer := clients.GetConsumer(bootstrapServers, topic, consumerGroup, consumerID, ssl)
		consumerName := fmt.Sprintf("%v-%v", consumerGroup, consumerID)

		fmt.Printf("[Consumer] Starting consumer %v\n", consumerName)

		go func() {
			for {
				m, err := consumer.ReadMessage(context.Background())
				if err != nil {
					wg.Done()
					break
				}

				atomic.AddUint64(&counter, 1)

				if verbose == true {
					fmt.Printf("[Key] %s | [Value] %s\n\n\n", m.Key, m.Value)
				}

				var multiple = counter % 100
				if multiple == 0 && counter != 0 {
					fmt.Printf("[Consumer] %v Messages retrived from topic %v by consumer group %s \n", counter, m.Topic, consumerGroup)
				}
			}
			wg.Done()
		}()

	}
	wg.Wait()
}

func createTopicBeforeTest(topic string, zookeeper string) {
	fmt.Printf("Creating topic %s\n", topic)
}

package clients

import (
	"kafka-stress/pkg/kafkadialer"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

// GetConsumer return a Kafka Consumer Client
func GetConsumer(bootstrapServers, topic, consumerGroup string, consumer int, ssl bool) *kafka.Reader {
	dialer := kafkadialer.GetDialer(ssl)

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        strings.Split(bootstrapServers, ","),
		Topic:          topic,
		GroupID:        consumerGroup,
		Dialer:         &dialer,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
	})
}

// GetProducer return a Kafka Producer Client
func GetProducer(bootstrapServers string, topic string, batchSize int, acks int, ssl bool, balancer string) *kafka.Writer {

	dialer := kafkadialer.GetDialer(ssl)

	var config kafka.WriterConfig

	config = kafka.WriterConfig{
		Brokers:      strings.Split(bootstrapServers, ","),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    batchSize,
		BatchTimeout: 2 * time.Second,
		RequiredAcks: acks,
		Dialer:       &dialer,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	switch balancer {
	case "hash":
		config.Balancer = &kafka.Hash{}
	case "murmur2":
		config.Balancer = &kafka.Murmur2Balancer{}
	case "crc32":
		config.Balancer = &kafka.CRC32Balancer{}				
	default: 
		config.Balancer = &kafka.Hash{}
	}	

	return kafka.NewWriter(config)

}

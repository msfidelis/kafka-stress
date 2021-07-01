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
		Brokers: strings.Split(bootstrapServers, ","),
		Topic:   topic,
		GroupID: consumerGroup,
		Dialer:  &dialer,
	})
}

// GetProducer return a Kafka Producer Client
func GetProducer(bootstrapServers string, topic string, batchSize int, acks int, ssl bool) *kafka.Writer {

	dialer := kafkadialer.GetDialer(ssl)

	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:      strings.Split(bootstrapServers, ","),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    batchSize,
		BatchTimeout: 2 * time.Second,
		// Balancer:     &kafka.Hash{},
		RequiredAcks: acks,
		Dialer:       &dialer,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	})

}

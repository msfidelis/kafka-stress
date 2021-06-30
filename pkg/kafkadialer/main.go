package kafkadialer

import (
	"crypto/tls"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

// GetDialer generate a kafka.Dialer instance configured
func GetDialer(ssl bool) kafka.Dialer {
	var dialer kafka.Dialer
	name, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	if ssl {
		dialer = kafka.Dialer{
			Timeout:   20 * time.Second,
			DualStack: true,
			ClientID:  name,
			TLS:       &tls.Config{},
		}
	} else {
		dialer = kafka.Dialer{
			Timeout:   20 * time.Second,
			DualStack: true,
			ClientID:  name,
		}
	}

	return dialer
}

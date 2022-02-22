# Kafka Stress - Stress Test Tool for Kafka Clusters, Producers and Consumers Tunning

<p>
  <a href="README.md" target="_blank">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" />
  </a>
  <a href="LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="/" target="_blank">
    <img alt="Build CI" src="https://github.com/msfidelis/kafka-stress/workflows/kafka-stress%20ci/badge.svg" />
  </a>  
  <a href="/" target="_blank">
    <img alt="Release" src="https://github.com/msfidelis/kafka-stress/workflows/release%20packages/badge.svg" />
  </a>
  <a href="https://twitter.com/fidelissauro" target="_blank">
    <img alt="Twitter: fidelissauro" src="https://img.shields.io/twitter/follow/fidelissauro.svg?style=social" />
  </a>  
</p>

# Introduction

> Kafka Stress is a simple CLI Tool to produce and consume events on Kafka topics in high throughput to tests and validate brokers, topics, partitions and consumers performances. 

# Installation 

### Docker 

```bash
docker pull fidelissauro/kafka-stress:latest
```

```bash
docker run --network host -it fidelissauro/kafka-stress:latest --bootstrap-servers 0.0.0.0:9092 --topic kafka-stress --test-mode consumer --consumers 6
```

### MacOS amd64

```bash
wget https://github.com/msfidelis/kafka-stress/releases/download/v0.0.5/kafka-stress_0.0.5_darwin_amd64 -O kafka-stress 
mv kafka-stress /usr/local/bin 
chmod +x /usr/local/bin/kafka-stress
```

### Linux amd64 

```bash
wget https://github.com/msfidelis/kafka-stress/releases/download/v0.0.5/kafka-stress_0.0.5_linux_amd64 -O kafka-stress 
mv kafka-stress /usr/local/bin 
chmod +x /usr/local/bin/kafka-stress
```

### Linux arm64 

```bash
wget https://github.com/msfidelis/kafka-stress/releases/download/v0.0.5/kafka-stress_0.0.5_linux_arm64 -O kafka-stress 
mv kafka-stress /usr/local/bin 
chmod +x /usr/local/bin/kafka-stress
```

### Freebsd amd64 

```bash
wget https://github.com/msfidelis/kafka-stress/releases/download/v0.0.5/kafka-stress_0.0.5_freebsd_amd64 -O kafka-stress 
mv kafka-stress /usr/local/bin 
chmod +x /usr/local/bin/kafka-stress
```

### Freebsd arm64 

```bash
wget https://github.com/msfidelis/kafka-stress/releases/download/v0.0.5/kafka-stress_0.0.5_freebsd_arm64 -O kafka-stress 
mv kafka-stress /usr/local/bin 
chmod +x /usr/local/bin/kafka-stress
```

# v0 Usage 

```bash
Usage of kafka-stress:
  -ack int
    	Required ACKs to produce messages (default 1)
  -batch-size int
    	Batch size for producer mode
  -bootstrap-servers string
    	Kafka Bootstrap Servers Broker Lists (default "0.0.0.0:9092")
  -consumer-group string
    	Consumer group name (default "kafka-stress")
  -consumers int
    	Number of consumers will be used in topic (default 1)
  -create-topic
    	Auto Create Topic?
  -events int
    	Numer of events will be created in topic (default 10000)
  -schema string
    	Schema
  -schema-registry string
    	Schema Registry URL (default "0.0.0.0:8081")
  -size int
    	Message size in bytes (default 62)
  -ssl-enabled
    	SSL Mode
  -test-mode string
    	Test Type; Ex producer;consumer. Default: producer (default "producer")
  -topic string
    	Kafka Stress Topics (default "kafka-stress")
  -zookeeper-servers string
    	Zookeeper Connection String (default "0.0.0.0:2181")
```

## Producer 

```bash
kafka-stress --bootstrap-servers localhost:29092 --events 30000 --topic kafka-stress
```

```bash
kafka-stress --bootstrap-servers localhost:29092 --events 10000 --topic kafka-stress

Sent 10000 messages to topic kafka-stress with 0 errors
Tests finished in 1.232463918s. Producer mean time 8113.83/s
```

### Customize ACK's 

Use `--ack` parameter to customize ack's quorum

```bash
kafka-stress --bootstrap-servers 0.0.0.0:9092 --events 20000 --topic tunning-3 --test-mode producer --ack 1
```

### Producer Format

You can produce random data in `string` and `json` format using `--format` flag

```bash
kafka-stress --bootstrap-servers 0.0.0.0:9092 --events 20000 --topic tunning-3 --test-mode producer --ack 1 --format json
```


## Consumer

kafka-stress --bootstrap-servers 0.0.0.0:9092 --topic kafka-stress --test-mode consumer --consumers 6
```

```bash
kafka-stress --bootstrap-servers 0.0.0.0:9092 --events 10000 --topic kafka-stress --test-mode consumer  --consumers 6

...
[Consumer 6] Message from consumer group kafka-stress at topic/partition/offset kafka-stress/1/0: 65b27cc0-1053-4fbc-b7a9-a40972fcaca6 = a124e95d-9226-4eb9-9169-411318bb6e4e
[Consumer 5] Message from consumer group kafka-stress at topic/partition/offset kafka-stress/2/0: 35fc905f-27f3-4b9f-81db-af89c1fa4c28 = 94f06092-589c-4030-8e3f-66a5a980123b
...
```

### Customize consumer-group name

Use `--consumer-group` to change customer group name used by workers. 

```bash
kafka-stress --bootstrap-servers 0.0.0.0:9092 --topic kafka-stress --test-mode consumer --consumer-group custom-consumer-group
```

## Authentication methods 

### SSL 

Use `--ssl-enabled` to enable ssl authentication. 

```bash
kafka-stress --bootstrap-servers 0.0.0.0:9092 --topic kafka-stress --test-mode consumer --ssl-enabled
```

## Roadmap 

Improve reports output
* Add execution time on reports - :check: 
* Add retry mechanism 
* Add message size options
* Add test header
* Add SCRAM authentication 
* Add SASL authetication 
* Add TLS authetication 
* Add IAM authentication for Amazon MSK 
* Add Unit Tests
* Add time based tests 
* Add Schema Registry / AVRO, JSON, PROTO Support

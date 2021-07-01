
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

Use `--ssl` to enable ssl authentication. 

```bash
kafka-stress --bootstrap-servers 0.0.0.0:9092 --topic kafka-stress --test-mode consumer --consumer-group custom-consumer-group
```

## Authentication methods 

### SSL 

Use `--ssl` to enable ssl authentication. 

```bash
kafka-stress --bootstrap-servers 0.0.0.0:9092 --topic kafka-stress --test-mode consumer --ssl
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
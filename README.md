
## v0 Usage 

```bash
Usage of kafka-stress:
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
  -test-mode string
    	Test Type; Ex producer;consumer. Default: producer (default "producer")
  -topic string
    	Kafka Stress Topics (default "kafka-stress")
  -zookeeper-servers string
    	Zookeeper Connection String (default "0.0.0.0:2181")
```

## Tests
1. Setup local stack 

```bash
docker-compose up --force-recreate
```

2. Create an Kafka Topic

```bash
docker-compose exec kafka  kafka-topics --create --topic kafka-stress --partitions 3 --replication-factor 1 --if-not-exists --zookeeper zookeeper:2181
```

3. Example usage 

### Producer Test

```bash
go run main.go --bootstrap-servers localhost:29092 --events 30000 --topic kafka-stress
```

```bash
❯ go run main.go --bootstrap-servers localhost:29092 --events 10000 --topic kafka-stress

Sent 10000 messages to topic kafka-stress with 0 errors
Tests finished in 1.232463918s. Producer mean time 8113.83/s
```

### Consumer Test

```bash
go run main.go --bootstrap-servers 0.0.0.0:9092 --topic kafka-stress --test-mode consumer --consumers 6
```

```bash
go run main.go --bootstrap-servers 0.0.0.0:9092 --events 10000 --topic kafka-stress --test-mode consumer  --consumers 6

...
[Consumer 6] Message from consumer group kafka-stress at topic/partition/offset kafka-stress/1/0: 65b27cc0-1053-4fbc-b7a9-a40972fcaca6 = a124e95d-9226-4eb9-9169-411318bb6e4e
[Consumer 5] Message from consumer group kafka-stress at topic/partition/offset kafka-stress/2/0: 35fc905f-27f3-4b9f-81db-af89c1fa4c28 = 94f06092-589c-4030-8e3f-66a5a980123b
...
```

## Roadmap 

Improve reports output
* Add execution time on reports - :check: 
* Add retry mechanism 
* Add message size options
* Add test header
* Add logs
* Add SCRAM authentication 
* Add SASL authetication 
* Add TLS authetication 
* Add IAM authentication for Amazon MSK 
* Add Unit Tests
* Add consume tests - :check: 
* Add time based tests 
* Add goreleaser pipeline 
* Add Docker image build
* Add Message Size Test 
* Add Schema Registry / AVRO, JSON, PROTO Support
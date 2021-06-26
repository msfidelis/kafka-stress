
## v0 Usage 

```bash
Usage of ./main:
  -bootstrap-servers string
    	Kafka Bootstrap Servers Broker Lists (default "0.0.0.0:9092")
  -create-topic
    	Auto Create Topic?
  -events int
    	Numer of events will be created in topic (default 10000)
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

```bash
go run main.go --bootstrap-servers localhost:29092 --events 30000 --topic kafka-stress
```

```bash
❯ go run main.go --bootstrap-servers localhost:29092 --events 10000 --topic fodase

Sent 10000 messages to topic fodase with 0 errors
Tests finished in 1.232463918s. Producer mean time 8113.83/s
```


## Roadmap 

* Improve reports output
* Add execution time on reports 
* Add retry mechanism 
* Add message size options
* Add test header
* Add logs
* Add SCRAM authentication 
* Add SASL authetication 
* Add TLS authetication 
* Add IAM authentication for Amazon MSK 
* Add Unit Tests
* Add consume tests 
* Add time based tests
* Add goreleaser pipeline 
* Add Docker image build
* Add Message Size Test 
* Add Schema Registry / AVRO, JSON, PROTO Support
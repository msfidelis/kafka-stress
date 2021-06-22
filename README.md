
## Tests
1. Setup local stack 

```
docker-compose up --force-recreate
```

2. Create an Kafka Topic

```
docker-compose exec kafka  kafka-topics --create --topic kafka-stress --partitions 3 --replication-factor 1 --if-not-exists --zookeeper zookeeper:2181
```

3. Example usage 

```
go run main.go --bootstrap-servers localhost:29092 --events 30000 --topic kafka-stress
```


## Roadmap 

* Improve reports output
* Add execution time on reports 
* Add retry mechanism 
* Add SCRAM authentication 
* Add SASL authetication 
* Add IAM authentication for Amazon MSK 
* Add Unit Tests
* Add consume tests 
* Add time based tests
* Add goreleaser pipeline 
* Add Docker image build
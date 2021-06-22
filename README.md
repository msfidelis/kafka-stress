```
docker-compose exec kafka  kafka-topics --create --topic kafka-stress --partitions 3 --replication-factor 1 --if-not-exists --zookeeper zookeeper:2181
```
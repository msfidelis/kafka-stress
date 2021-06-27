
# Schema Registry

## Utils

[Json to AVRO Converter](https://toolslick.com/generation/metadata/avro-schema-from-json)

## List Schemas 

```bash
curl -X GET http://0.0.0.0:8081/subjects
```

## AVRO

Create and Schema in AVRO format 

```
curl http://0.0.0.0:8081/subjects/supis/versions -X POST \
-H "Content-Type: application/vnd.schemaregistry.v1+json" \
-d '
{
   "schema":"{\n      \"name\":\"Example\",\n      \"type\":\"record\",\n      \"namespace\":\"com.acme.avro\",\n      \"fields\":[\n         {\n            \"name\":\"name\",\n            \"type\":\"string\"\n         },\n         {\n            \"name\":\"age\",\n            \"type\":\"int\"\n         }\n      ]\n   }"
}
'
```



# Consumer 

## Algoritms 

Producing 15000 events in 3 partitions topic

### Hash 

```bash
❯ go run main.go --bootstrap-servers 0.0.0.0:9092 --events 15000 --topic brabo --test-mode producer  --consumers 3
Sent 15000 messages to topic brabo with 0 errors
Tests finished in 1.406677516s. Producer mean time 10663.42/s
```

### LeastBytes

```bash
❯ go run main.go --bootstrap-servers 0.0.0.0:9092 --events 15000 --topic brabo --test-mode producer  --consumers 3
Sent 15000 messages to topic brabo with 0 errors
Tests finished in 885.728755ms. Producer mean time 16935.21/s
```
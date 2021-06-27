
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


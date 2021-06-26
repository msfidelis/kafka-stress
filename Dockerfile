FROM golang:1.15 AS builder

WORKDIR $GOPATH/src/kafka-stress

COPY . ./

RUN go get -u

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kafka-stress .


FROM alpine:latest

COPY --from=builder /go/src/kafka-stress/kafka-stress ./


ENTRYPOINT ["./kafka-stress"]
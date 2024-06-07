FROM golang:1.22 AS builder

WORKDIR $GOPATH/src/kafka-stress

COPY . ./

RUN go get -u

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kafka-stress .


FROM cgr.dev/chainguard/wolfi-base:latest

COPY --from=builder /go/src/kafka-stress/kafka-stress ./


ENTRYPOINT ["./kafka-stress"]
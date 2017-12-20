FROM golang:1.8.1 as builder

WORKDIR /go/src/github.com/net-one/gcpexa/services/instance-client

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM alpine:latest


RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/net-one/gcpexa/services/instance-client .

CMD ["./instance-client"]

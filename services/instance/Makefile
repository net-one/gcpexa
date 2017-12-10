build:  
    protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/gcpexa/services/instance \
      proto/createInstance.proto

    GOOS=linux GOARCH=amd64 go build
    docker build -t instance .

run:  
    docker run -p 50051:50051 instance
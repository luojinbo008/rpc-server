package main

/*
import (
	"../thrift-go-gen/leader"
	"git.apache.org.bak/thrift.git/lib/go/thrift"
	"fmt"
	"context"
	"log"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

type leaderThrift struct {}

func (fdi *leaderThrift) Register(ctx context.Context, data *leader.Leader) (x bool, err error){

	fmt.Println("11111111111111")
	return true, nil
}

func main() {
	handler := &leaderThrift{}
	processor := leader.NewLeaderThriftProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket(HOST + ":" + PORT)

	if err != nil {
		log.Fatalln("Error:", err)
	}

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", HOST + ":" + PORT)
	server.Serve()
}*/

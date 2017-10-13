package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"context"
	"../server/thrift/thrift-go-gen/leader"
	"log"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func main()  {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, err := transportFactory.GetTransport(tSocket)

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := leader.NewLeaderThriftClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST + ":" + PORT)
	}
	defer transport.Close()


	data := leader.Leader{
		1,
		1,
		"test_123",
		"罗锦波",
		"330682198808193416",
		18667188875,
		1,
		"测试的",
		}
	d, err := client.Register(context.Background(), &data)
	fmt.Println(d)
}
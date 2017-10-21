
package main

import (
	"server/thriftRPC/guide"
	"server/handler"
	"github.com/luojinbo008/bfgo"
	"log"
)


func main() {
	processor := guide.NewGuideThriftProcessor(new(handler.GuideThrift))

	err := bfgo.Init("thrift.toml", processor)

	if err != nil {
		log.Fatal(err)
	}
	err = bfgo.Run()
	if err != nil {
		log.Fatal(err)
	}
}

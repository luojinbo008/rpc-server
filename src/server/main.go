package main

import (
	"github.com/luojinbo008/thrift_manage"
	"server/login"
)


func main() {
	thrift_manage.Run(
		login.Module,
	)
}
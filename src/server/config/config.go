package conf

import (
	"log"
)

type ConfigSlice struct {
	Database	[]Db_Config
}

var Configs ConfigSlice

var (
	// log conf
	LogFlag = log.LstdFlags
)
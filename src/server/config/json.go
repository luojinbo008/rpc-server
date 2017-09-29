package conf

import (
	"encoding/json"
	"log"
	"io/ioutil"
)

type Db_Config struct {
	Host		string
	Username	string
	Password	string
	Dbname		string
	Charset		string
}

func init() {
	data, err := ioutil.ReadFile("config/config.json")

	if err != nil {
		log.Fatal("%v", err)
	}

	err = json.Unmarshal(data, &Configs)

	if err != nil {
		log.Fatal("%v", err)
	}
}
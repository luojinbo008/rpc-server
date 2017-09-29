package container

import(
	lconf "./config"
	"fmt"
	//database/sql"
)

type db_element struct {

}

func (db *db_element) Run() {

	for _, config := range lconf.Configs.Database {
		fmt.Printf("%v", config)
	}

/*	DB_center, _ = sql.Open("mysql", "root:@tcp(192.168.56.101:3306)/newwechat?charset=utf8")
	DB_center.SetMaxOpenConns(20)
	DB_center.SetMaxIdleConns(10)
	DB_center.Ping()*/
}

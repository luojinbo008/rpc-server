package mysql

import (
	"github.com/luojinbo008/bfgo/app"
	"github.com/jinzhu/gorm"
	"server/dao/mysql"
)

type DBNameModel int

func NewDBNameDao() *mysql.DBNameDao {
	userModel := new(DBNameModel)
	db := app.UseModel("mysql", userModel.ConnName(), true).(*gorm.DB)
	return &mysql.DBNameDao{DB:db}
}

func (p *DBNameModel) ConnName() string {
	return "guide"
}
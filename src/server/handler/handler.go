package handler

import (
	"server/thriftRPC/guide"
	"server/model/mysql"
	dmysql "server/dao/mysql"
	"context"
	"fmt"
)


type GuideThrift struct {

}

func (fdi *GuideThrift) Register(ctx context.Context, data *guide.Guide) (bool, error) {

	mDB := mysql.NewDBNameDao()

	data1, err := mDB.AddGuideUser(dmysql.GuideUser{
		daya.app_id,
	})

	if err != nil{
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", data1)
	return true, nil
}
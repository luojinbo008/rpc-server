package mysql

import (
	"github.com/jinzhu/gorm"
)

type DBNameDao struct {
	DB *gorm.DB
}

type GuideUser struct {
	ID 				uint		`gorm:"primary_key:id,not null"`
	AppID			uint		`gorm:"column:appid,not null"`
	NickName		string		`gorm:"column:nickname,not null"`
	RealName		string		`gorm:"column:realname,not null"`
	IdCard			string		`gorm:"column:idcard,not null"`
	Phone			string		`gorm:"column:phone,not null"`
	Status 			uint8		`gorm:"column:status,not null"`
	Birthday		uint64		`gorm:"column:birthday"`
	Memo 			string		`gorm:"column:memo"`
	AddTime			uint64		`gorm:"column:addtime,not null"`
}

func(p *DBNameDao) AddGuideUser(userGuide GuideUser) (id uint, err error) {
	p.DB.Table("guideUser").Create(&userGuide)
	return 1, nil
}

func(p *DBNameDao) SelectDevicenuByUseridSource(userID int, source int) (userDevice GuideUser, err error) {
	p.DB.Table("guideUser").Where("id = ?", userID).Where("source = ?", source).Find(&userDevice)
	return
}

/*
CREATE TABLE `guideUser` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `appid` int(11) NOT NULL,
  `nickname` varchar(100) NOT NULL,
  `realname` varchar(20) NOT NULL,
  `idcard` varchar(50) NOT NULL,
  `phone` varchar(50) NOT NULL,
  `status` tinyint(3) NOT NULL DEFAULT '0',
  `birthday` date NOT NULL,
  `memo` varchar(200) NOT NULL,
  `addTime` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `u-nickname` (`nickname`,`appid`) USING BTREE,
  UNIQUE KEY `u-phone` (`phone`,`appid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

 */
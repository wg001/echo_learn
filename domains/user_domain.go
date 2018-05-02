package domains

import (
	"echo_learn/dto"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"echo_learn/utils"
)

const TABLE_USER = "user" //用户表

var default_userorm dto.UserORM

func GetUserInfo(id uint64, db *gorm.DB) dto.UserORM {
	userORM := dto.UserORM{}
	db.Table(TABLE_USER).Where("id=?", id).Find(&userORM)
	if userORM != default_userorm {
		(&userORM).CreateTime=utils.FormatTimeToLocal(userORM.CreateTime)
		(&userORM).UpdateTime=utils.FormatTimeToLocal(userORM.UpdateTime)
	}
	log.Infof("dddddd--%v\n",userORM)
	return userORM
}

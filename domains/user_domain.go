package domains

import (
	"echo_learn/dto"
	"echo_learn/utils"
	"echo_learn/errors"
	"github.com/sirupsen/logrus"
)

const TABLE_USER = "user" //用户表

var default_userorm dto.UserORM

func GetUserInfo(id uint64) (dto.UserORM, error) {
	userobj := dto.UserORM{}
	dbobj, err := dto.GetDbStruct()
	if err != nil {
		return userobj, errors.NewCommonError(errors.ERROR_TYPE_DBCONNECTION, err.Error())
	}
	db := dbobj.GetDBObj()
	db.Table(TABLE_USER).Where("id=?", id).Find(&userobj)
	if userobj != default_userorm {
		(&userobj).CreateTime = utils.FormatTimeToLocal(userobj.CreateTime)
		(&userobj).UpdateTime = utils.FormatTimeToLocal(userobj.UpdateTime)
	}
	logrus.Info("ddddd--%v\n", userobj)
	return userobj, nil
}

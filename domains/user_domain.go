package domains

import (
	"echo_learn/dto"
	"echo_learn/utils"
	"echo_learn/errors"
	"github.com/sirupsen/logrus"
	"fmt"
	"time"
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

func GetUserByCondition(condition dto.UserORM) (dto.UserORM, error) {
	dbobj, err := dto.GetDbStruct()
	if err != nil {
		return dto.UserORM{}, errors.NewCommonError(errors.ERROR_TYPE_DBCONNECTION, err.Error())
	}
	db := dbobj.GetDBObj()
	searchUser:=dto.UserORM{}
	db.Table(TABLE_USER).Where(&condition).Find(&searchUser)
	return searchUser,nil
}

func SaveUser(userInfo dto.Register) error {
	userOrm := &dto.UserORM{}
	userOrm.Phone = userInfo.Phone
	userOrm.CreateTime = time.Now()
	//userOrm.UpdateTime = time.Date(2017,time.December,14,12,12,45,2e8, time.UTC)
	userOrm.UpdateTime, _ = time.Parse("2006-01-02 15:04:05", "2014-11-11 11:11:12")
	dbobj, err := dto.GetDbStruct()
	if err != nil {
		return errors.NewCommonError(errors.ERROR_TYPE_DBCONNECTION, err.Error())
	}
	db := dbobj.GetDBObj()
	affectedrows := db.Table(TABLE_USER).Create(userOrm).RowsAffected
	fmt.Println(affectedrows)
	return nil
}

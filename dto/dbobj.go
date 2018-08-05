package dto

import (
	"sync"
	"github.com/jinzhu/gorm"
	"log"
	"github.com/sirupsen/logrus"
	"github.com/pkg/errors"
)

type dbobj struct {
	//同步锁
	Lock sync.Mutex
	//是否初始化
	Initialize bool
	//数据库连接对象
	DbObject *gorm.DB
}

type DBObj interface {
	SetDB(db *gorm.DB) error
	GetDBObj() *gorm.DB
}

var DBObjInstance *dbobj

func GetDbStruct() (DBObj,error){
	if DBObjInstance==nil{
		dbDSN := globalConf.Mysql.UserName+":"+globalConf.Mysql.PassWord+"@tcp("+globalConf.Mysql.Host+":"+globalConf.Mysql.Port+")/"+globalConf.Mysql.DataBase+"?charset="+
			globalConf.Mysql.Charset+"&parseTime="+globalConf.Mysql.ParseTime
		db,err:=gorm.Open("mysql",dbDSN)
		if err!=nil{
			logrus.Error("数据库链接失败，错误信息：",err.Error())
			return nil,errors.New("数据库连接失败啦")
		}
		DBObjInstance = &dbobj{Initialize:true,DbObject:db}
		log.Printf("-----%v\n",DBObjInstance)
		return DBObjInstance,nil
	}
	return DBObjInstance,nil
}

func (obj *dbobj)GetDBObj() *gorm.DB {
	return obj.DbObject
}

func (obj *dbobj)SetDB(db *gorm.DB) error {
	obj.Lock.Lock()
	defer obj.Lock.Unlock()
	if !obj.Initialize{
		obj.DbObject = db
		obj.Initialize=true
	}
	return nil
}
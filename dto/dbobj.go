package dto

import (
	"sync"
	"github.com/jinzhu/gorm"
	"log"
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

func GetDbStruct() DBObj{
	if DBObjInstance==nil{
		dbDSN := globalConf.Mysql.UserName+":"+globalConf.Mysql.PassWord+"@tcp("+globalConf.Mysql.Host+":"+globalConf.Mysql.Port+")/"+globalConf.Mysql.DataBase+"?charset="+
			globalConf.Mysql.Charset+"&parseTime="+globalConf.Mysql.ParseTime
		db,err:=gorm.Open("mysql",dbDSN)
		if err!=nil{
			log.Fatalf("连接数据库出错:%v",err)
		}
		DBObjInstance = &dbobj{Initialize:true,DbObject:db}
		log.Printf("-----%v\n",DBObjInstance)
		return DBObjInstance
	}
	return DBObjInstance
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
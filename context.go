package orm

import (
	"database/sql"
	"sync"
)

// Orm Context
type OrmContext struct {
	conn *sql.DB
}
// 'OrmContext'指针变量
var ormContext *OrmContext
// 同步控制变量
var ormContextLock sync.Once
// 单例'OrmContext'
func singleOrmContext() OrmContext {
	ormContextLock.Do(func() {
		ormContext = new(OrmContext)
		driverName := "mysql"
		dataSourceName := "umesample:umePW123!!@tcp(114.115.185.91:3306)/umesample?charset=utf8&loc=Asia%2FShanghai&parseTime=true"
		db, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			panic(err)
		}
		//defer db.Close()
		ormContext.conn = db

	})
	return *ormContext
}

//// 初始化全局上下文实例
//func (owner *OrmContext) Open() *OrmContext {
//	return singleOrmContext()
//}

// 释放上下文
func (owner *OrmContext) Close() {
	owner.conn.Close()
}

// 获取数据库访问实例
func (owner *OrmContext) DB()  *sql.DB {
	return owner.conn
}





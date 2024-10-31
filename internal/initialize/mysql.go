package initialize

import (
	"fmt"
	"time"

	"example.com/m/global"
	"example.com/m/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false, // on/off transaction in mysql that help to improve performance
	})
	checkErrorPanic(err, "Mysql init Error")
	global.Mdb = db

	// set pool
	SetPool()
	migrateTables()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("Mysql error %s: ", err)
	}

	sqlDB.SetMaxIdleConns(m.MaxIdleConns)                       // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)                       // SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnsMaxLifeTime)) // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Println("Migrate table error: ", err)
	}
}

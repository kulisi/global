package internal

import (
	"github.com/kulisi/global/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

func Mssql(m config.Mssql) *gorm.DB {
	//m := ApiConfig.Mssql
	if m.Dbname == "" {
		log.Println("db-name is nil")
		return nil
	}
	conf := sqlserver.Config{
		DSN:               m.Dsn(),
		DefaultStringSize: 255,
	}

	if db, err := gorm.Open(sqlserver.New(conf), Config(m.GeneralDB)); err != nil {
		log.Println(err.Error())
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

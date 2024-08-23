package internal

import (
	"github.com/kulisi/global/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Mssql(m config.Mssql) *gorm.DB {
	//m := ApiConfig.Mssql
	if m.Dbname == "" {
		return nil
	}
	conf := sqlserver.Config{
		DSN:               m.Dsn(),
		DefaultStringSize: 255,
	}

	if db, err := gorm.Open(sqlserver.New(conf), Config(m.GeneralDB)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

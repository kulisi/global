package internal

import (
	"github.com/kulisi/global/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql(m config.Mysql) *gorm.DB {
	//m := global.ApiConfig.Mysql
	if m.Dbname == "" {
		return nil
	}
	conf := mysql.Config{
		DSN:               m.Dsn(),
		DefaultStringSize: 255,
	}

	if db, err := gorm.Open(mysql.New(conf), Config(m.GeneralDB)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

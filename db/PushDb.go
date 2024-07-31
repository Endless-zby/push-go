package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"push-go/config"
	"push-go/entity"
)

var PushDb *gorm.DB

func InitDb(cfg *config.DatabaseConfig) {
	var err error
	var directory gorm.Dialector

	switch cfg.Type {
	case "sqlite":
		directory = sqlite.Open(cfg.Sqlite.File)
	case "mysql":
		connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Dbname)
		directory = mysql.Open(connectionString)
	default:
		panic(fmt.Sprintf("Unsupported database type: %s", cfg.Type))
	}

	PushDb, err = gorm.Open(directory, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return
	}

	// orm
	PushDb.AutoMigrate(&entity.ClientDrive{})
	PushDb.AutoMigrate(&entity.MessageHistory{})
}

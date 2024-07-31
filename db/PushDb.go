package db

import (
	"flag"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"push-go/entity"
)

var PushDb *gorm.DB

func init() {
	// Define command-line flags
	dbType := flag.String("db", "mysql", "Database type: sqlite or mysql")
	flag.Parse()

	var err error
	var directory gorm.Dialector

	switch *dbType {
	case "sqlite":
		directory = sqlite.Open("push.db")
	case "mysql":
		mysqlDSN := "byzhao:zby123456@tcp(192.168.192.36:3306)/byzhao?charset=utf8mb4&parseTime=True&loc=Local"
		directory = mysql.Open(mysqlDSN)
	default:
		panic(fmt.Sprintf("Unsupported database type: %s", *dbType))
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

package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"push-go/entity"
)

var PushDb *gorm.DB

func init() {
	var err error
	PushDb, err = gorm.Open(sqlite.Open("push.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	PushDb.AutoMigrate(&entity.ClientDrive{})
	PushDb.AutoMigrate(&entity.MessageHistory{})
}

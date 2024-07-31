package entity

import (
	"gorm.io/gorm"
	"time"
)

type ClientDrive struct {
	//gorm.Model	添加一些默认的属性  例如：
	//-	  `ID` ：每个记录的唯一标识符（主键）。
	//-   `CreatedAt` ：在创建记录时自动设置为当前时间。
	//-   `UpdatedAt`：每当记录更新时，自动更新为当前时间。
	//-   `DeletedAt`：用于软删除（将记录标记为已删除，而实际上并未从数据库中删除）。
	Id       uint   `gorm:"column:id;primaryKey" json:"id"`
	ClientId string `gorm:"column:client_id;unique;not null;size:64" json:"clientId"`
	DriveId  string `gorm:"column:drive_id;not null;size:255" json:"driveId"`
}

func (ClientDrive) TableName() string {
	return "token"
}

type MessageHistory struct {
	//gorm.Model
	Id        uint   `gorm:"column:id;primaryKey" json:"id"`
	ClientId  string `gorm:"column:client_id;not null;size:64" json:"clientId"`
	DriveId   string `gorm:"column:drive_id;not null;size:255" json:"driveId"`
	Title     string `gorm:"column:title" json:"title"`
	Body      string `gorm:"column:body;not null" json:"body"`
	Timestamp uint   `gorm:"column:timestamp" json:"timestamp"`
}

func (MessageHistory) TableName() string {
	return "history"
}

func (m *MessageHistory) BeforeCreate(tx *gorm.DB) (err error) {
	m.Timestamp = uint(time.Now().Unix())
	return
}

func (m *MessageHistory) BeforeUpdate(tx *gorm.DB) (err error) {
	m.Timestamp = uint(time.Now().Unix())
	return
}

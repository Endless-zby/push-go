package entity

import (
	"gorm.io/gorm"
	"time"
)

type ClientDrive struct {
	//gorm.Model	添加一些默认的属性  例如：ID CreatedAt UpdatedAt DeletedAt   总结：没啥用 鸡肋
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

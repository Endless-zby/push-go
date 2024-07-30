package entity

type ClientDrive struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	ClientId string `json:"clientId"`
	DriveId  string `json:"driveId"`
}

type MessageHistory struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	ClientId  string `json:"clientId"`
	DriveId   string `json:"driveId"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Timestamp uint   `json:"timestamp"`
}

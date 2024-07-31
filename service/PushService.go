package service

import (
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"
	"log"
	"push-go/apns"
	"push-go/db"
	"push-go/entity"
)

func PushAndSave(message entity.MessageHistory) {
	notification := &apns2.Notification{}
	notification.DeviceToken = message.DriveId
	notification.Topic = "me.fin.bark"

	payloadStr := payload.NewPayload().AlertBody(message.Body).AlertTitle(message.Title)

	notification.Payload = payloadStr
	_, err := apns.Client.Push(notification) // 不用关心发送结果
	if err != nil {
		log.Printf("消息发送失败 err: %s", err.Error())
	}
	// todo 发送消息并存储
	db.PushDb.Save(&message)
}

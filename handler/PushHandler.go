package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"push-go/db"
	"push-go/entity"
	"push-go/service"
	"strconv"
)

// 消息发送

func SendGet(c *gin.Context) {
	var messageHistory entity.MessageHistory
	messageHistory.ClientId = c.Query("clientId")
	messageHistory.Title = c.Query("title")
	messageHistory.Body = c.Query("body")
	// 根据clientId查询driveToken
	clientDrive, errs := db.FindByClientId(messageHistory.ClientId)
	if errs != nil {
		errStr := fmt.Sprintf("failed to find client drive for clientId %s: %v", messageHistory.ClientId, errs)
		c.JSON(http.StatusOK, entity.IsFailMessage(errStr))
		return
	}
	messageHistory.DriveId = clientDrive.DriveId
	log.Printf("SendGet: title:%s, body:%s, clientId:%s driver:%s", messageHistory.Title, messageHistory.Body, messageHistory.ClientId, messageHistory.DriveId)

	go service.PushAndSave(messageHistory)

	c.JSON(http.StatusOK, entity.IsSuccess())
}
func SendPost(c *gin.Context) {
	var messageHistory entity.MessageHistory
	err := c.ShouldBindJSON(&messageHistory)
	if err != nil {
		c.JSON(http.StatusOK, entity.IsFailNoMessage())
		return
	}
	// 根据clientId查询driveToken
	clientDrive, err := db.FindByClientId(messageHistory.ClientId)
	if err != nil {
		errStr := fmt.Sprintf("failed to find client drive for clientId %s: %v", messageHistory.ClientId, err)
		c.JSON(http.StatusOK, entity.IsFailMessage(errStr))
		return
	}
	messageHistory.DriveId = clientDrive.DriveId
	log.Printf("SendPost: title:%s, body:%s, clientId:%s driver:%s", messageHistory.Title, messageHistory.Body, messageHistory.ClientId, messageHistory.DriveId)

	go service.PushAndSave(messageHistory)

	c.JSON(http.StatusOK, entity.IsSuccess())
}

func SendParam(c *gin.Context) {
	var messageHistory entity.MessageHistory
	messageHistory.ClientId = c.Param("clientId")
	messageHistory.Title = c.Param("title")
	messageHistory.Body = c.Param("body")
	// 根据clientId查询driveToken
	clientDrive, err := db.FindByClientId(messageHistory.ClientId)
	if err != nil {
		errStr := fmt.Sprintf("failed to find client drive for clientId %s: %v", messageHistory.ClientId, err)
		c.JSON(http.StatusOK, entity.IsFailMessage(errStr))
		return
	}
	messageHistory.DriveId = clientDrive.DriveId
	log.Printf("SendPost: title:%s, body:%s, clientId:%s driver:%s", messageHistory.Title, messageHistory.Body, messageHistory.ClientId, messageHistory.DriveId)

	go service.PushAndSave(messageHistory)

	c.JSON(http.StatusOK, entity.IsSuccess())
}

// 初始化token

func SaveDriveGet(c *gin.Context) {
	clientId := c.Query("clientId")
	driveId := c.Query("driveId")
	clientDrive := entity.ClientDrive{
		ClientId: clientId,
		DriveId:  driveId,
	}
	db.PushDb.Save(&clientDrive)
	log.Printf("SendGet: clientId:%s, driveId:%s", clientId, driveId)
	// 保存
	c.JSON(http.StatusOK, entity.IsSuccess())
}

func SaveDrivePost(c *gin.Context) {
	var clientDrive entity.ClientDrive
	err := c.ShouldBindJSON(&clientDrive)
	if err != nil {
		c.JSON(http.StatusOK, entity.IsFailNoMessage())
		return
	}
	db.PushDb.Save(&clientDrive)
	log.Printf("SendGet: clientId:%s, driveId:%s", clientDrive.ClientId, clientDrive.DriveId)
	// 保存
	c.JSON(http.StatusOK, entity.IsSuccess())
}

// 查询Drivetoken

func DriveGet(c *gin.Context) {
	clientId := c.Query("clientId")
	if clientId == "" {
		c.JSON(http.StatusOK, entity.IsFailMessage("异常参数"))
		return
	}
	log.Printf("SendGet: clientId:%s", clientId)
	clientDrive, err := db.FindByClientId(clientId)
	if err != nil {
		errStr := fmt.Sprintf("failed to find client drive for clientId %s: %v", clientId, err)
		c.JSON(http.StatusOK, entity.IsFailMessage(errStr))
		return
	}
	c.JSON(http.StatusOK, entity.IsSuccessData(clientDrive))
}

// 历史消息

func HistoryGet(c *gin.Context) {
	clientId := c.Query("clientId")
	if clientId == "" {
		c.JSON(http.StatusOK, entity.IsFailMessage("异常参数"))
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	log.Printf("SendGet: clientId:%s, page:%d, size:%d", clientId, page, size)
	limitData, err := db.FindAllByClientIdLimit(clientId, page, size)
	if err != nil {
		errStr := fmt.Sprintf("failed to find client drive for clientId %s: %v", clientId, err)
		c.JSON(http.StatusOK, entity.IsFailMessage(errStr))
		return
	}
	c.JSON(http.StatusOK, entity.IsSuccessData(limitData))
}

package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"push-go/db"
	"push-go/entity"
	"strconv"
)

// 消息发送

func SendGet(c *gin.Context) {
	//var messageHistoryDto MessageHistoryDto
	//if err := c.ShouldBindQuery(&messageHistoryDto); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"result": "序列化失败",
	//	})
	//	return
	//}
	title := c.Query("title")
	body := c.Query("body")
	clientId := c.Query("clientId")
	// 根据clientId查询driveToken
	log.Printf("SendGet: title:%s, body:%s, clientId:%s", title, body, clientId)
	// 发送消息
	c.JSON(http.StatusOK, entity.IsSuccess())
}
func SendPost(c *gin.Context) {
	var messageHistory entity.MessageHistory
	err := c.ShouldBind(&messageHistory)
	if err != nil {
		c.JSON(http.StatusOK, entity.IsFailNoMessage())
		return
	}

	// 根据clientId查询driveToken
	log.Printf("SendGet: title:%s, body:%s, clientId:%s", messageHistory.Title, messageHistory.Body, messageHistory.ClientId)
	// 发送消息

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
	err := c.ShouldBind(&clientDrive)
	if err != nil {
		c.JSON(http.StatusOK, entity.IsFailNoMessage())
		return
	}

	// 根据clientId查询driveToken
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
	clientDrive := entity.ClientDrive{ClientId: "sddd", DriveId: "dddd"}
	c.JSON(http.StatusOK, entity.IsSuccessData(clientDrive))
}

// 历史消息

func HistoryGet(c *gin.Context) {
	clientId := c.Query("clientId")
	if clientId == "" {
		c.JSON(http.StatusOK, entity.IsFailMessage("异常参数"))
		return
	}
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	sizeStr := c.Query("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < 1 {
		size = 1
	}
	log.Printf("SendGet: clientId:%s, page:%d, size:%d", clientId, page, size)
	messageHistoryDto := []entity.MessageHistory{{ClientId: "sddd", DriveId: "dddd"}, {ClientId: "sddd", DriveId: "dddd"}}
	c.JSON(http.StatusOK, entity.IsSuccessData(messageHistoryDto))
}

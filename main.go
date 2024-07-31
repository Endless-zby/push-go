package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"push-go/apns"
	"push-go/config"
	"push-go/db"
	"push-go/handler"
)

// --------------热加载调试使用和说明--------------------
/**
Python 的 Flask框架，有 debug 模式，启动时传入 debug=True 就可以热加载(Hot Reload, Live Reload)了。即更改源码，保存后，自动触发更新，浏览器上刷新即可。免去了杀进程、重新启动之苦。
Gin 原生不支持，但有很多额外的库可以支持。例如
- github.com/codegangsta/gin
- github.com/pilu/fresh
在这，我们尝试采用 github.com/pilu/fresh
go get -v -u github.com/pilu/fresh

安装好之后，记得先go build一下，mod中会生成执行文件fresh.exe，
只需要将go run main.go命令换成fresh即可,直接执行fresh: .\fresh
fresh会启动main文件并监控go程序的改动

以后 的每次更改源文件，代码将自动重新编译(Auto Compile)。

*/

func main() {
	configPath := flag.String("config", "config/config.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	db.InitDb(&cfg.Database) // 初始化数据库
	apns.Init(&cfg.Apns)     // 初始化apns连接
	server := gin.Default()
	server.Use(handler.Logger())
	//server.Use(config.GlobalMiddleWare)
	king := server.Group("/king", handler.GlobalMiddleWare)
	{
		king.GET("/send", handler.SendGet)
		king.POST("/send", handler.SendPost)
		king.GET("/send/:clientId/:title/:body", handler.SendParam)

		king.GET("/saveDrive", handler.SaveDriveGet)
		king.POST("/saveDrive", handler.SaveDrivePost)

		king.GET("/getDrive", handler.DriveGet)

		king.GET("/getHistory", handler.HistoryGet)
	}
	server.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}

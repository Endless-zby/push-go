package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"push-go/config"
	"push-go/controller"
	"time"
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
	// 首先，我们使用了gin.Default()生成了一个实例，这个实例即 WSGI(Web服务器网关接口) 应用程序。
	server := gin.Default()
	// 这是两种中间件的加载方式
	server.Use(config.Logger())
	//server.Use(config.GlobalMiddleWare)

	king := server.Group("/king", config.GlobalMiddleWare)
	{
		king.GET("/send", controller.SendGet)
		king.POST("/send", controller.SendPost)

		king.GET("/saveDrive", controller.SaveDriveGet)
		king.POST("/saveDrive", controller.SaveDrivePost)

		king.GET("/getDrive", controller.DriveGet)

		king.GET("/getHistory", controller.HistoryGet)
	}

	// 接下来，我们使用r.Get("/", ...)声明了一个路由，告诉 Gin 什么样的URL 能触发传入的函数，这个函数返回我们想要显示在用户浏览器中的信息。
	server.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello, Gin")
	})
	server.GET("/hello", func(context *gin.Context) {
		for index, value := range context.Request.URL.Query() {
			fmt.Println(index, value)
		}
		name := context.Query("name")
		if name == "" {
			name = "byzhao"
		}
		context.JSON(http.StatusOK, gin.H{
			"hello": name,
		})
	})

	// 路由带有参数
	server.GET("/hello/:name", func(context *gin.Context) {
		context.String(http.StatusOK, context.Param("name"))
	})

	// 获取query参数
	server.GET("/hello/name/:name", func(context *gin.Context) {
		name := context.Param("name")
		id := context.Query("id")
		age := context.DefaultQuery("age", "20")
		context.String(http.StatusOK, "hello %s, id:%s, age:%s", name, id, age)
	})

	// post 获取参数
	server.POST("/insert", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, map[string]string{
			"username": username,
			"password": password,
		})

	})

	server.POST("", func(context *gin.Context) {
		var student Student
		err := context.ShouldBindQuery(&student)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"dfd": "fff",
			})
			return
		}

	})

	// query 和 post 混合参数
	server.POST("/update", func(context *gin.Context) {
		id := context.Query("id")
		username := context.PostForm("username")
		password := context.DefaultPostForm("password", "zby123456")
		context.JSON(http.StatusOK, gin.H{
			"id":       id,
			"username": username,
			"password": password,
		})
	})

	// Map参数
	server.POST("/map", func(context *gin.Context) {
		ids := context.QueryMap("ids")
		names := context.PostFormMap("names")
		context.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})
	// 重定向
	server.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/hello")
	})

	server.GET("/goindex", func(context *gin.Context) {
		context.Request.URL.Path = "/hello"
		server.HandleContext(context)
	})

	// 绑定post请求体中的json数据
	server.POST("/insert/student", func(context *gin.Context) {
		var student Student
		err := context.ShouldBindJSON(&student)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"result": err.Error(),
			})
			return
		}
		var ca Car = &Maserati{}
		flag, _ := ca.getFlag()
		context.JSON(http.StatusOK, gin.H{
			"result": "OK",
			"data":   student,
			"name":   student.getName("zzzz"),
			"flag":   flag,
			"horn":   ca.honkTheHorn(),
		})
	})
	// 绑定query中的参数
	server.GET("/insert/car", func(context *gin.Context) {
		var ma Maserati

		if err := context.ShouldBindQuery(&ma); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"result": "序列化失败",
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"result": "OK",
			"data":   ma,
		})
	})

	// -------------------------------分组路由-------------------------------

	DefaultHandle := func(context *gin.Context) {
		fmt.Println("进入到业务执行器")
		all, _ := context.Get("all")
		fmt.Println("从中间件取出的数据为：", all)
		context.JSON(http.StatusOK, gin.H{
			"result": context.FullPath(),
		})
		fmt.Println("退出业务执行器")
	}
	// 局部中间件是针对某一个路由而执行的，不是所有路由都执行的
	// 如果存在全局中间件，则执行流程为：先执行全局中间件，再执行局部中间件
	// 中间件按照加载顺序执行，依次进入，逐层退出
	v1 := server.Group("/v1", config.GlobalMiddleWare, DefaultHandle)
	{
		v1.GET("/posts")
		v1.GET("/series")
	}
	v2 := server.Group("/v2", config.NoNextFunction, DefaultHandle)
	{
		v2.GET("/posts")
		v2.GET("/series")
	}

	// 最后用 r.Run()函数来让应用运行在本地服务器上，默认监听端口是 _8080_，可以传入参数设置端口，例如r.Run(":9999")即运行在 _9999_端口。
	server.Run()
}

// -------------------------------中间件-------------------------------

// 所有的请求都会经过这个中间件
func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		// 给context 实例设置一个值
		context.Set("gintime", t)
		// 请求前
		context.Next()
		// 请求后
		lateny := time.Since(t)
		log.Print(lateny)
	}
}

type Student struct {
	Name  string `json:"name" binding:"required"`
	Grade int    `json:"grade" binding:"required"`
	Age   int    `json:"age"`
}

type ApiResponse struct {
	Code       int         `json:"code"`
	ErrMessage string      `json:"errMessage"`
	Data       interface{} `json:"data"`
}

func (stu *Student) getName(name string) string {
	return stu.Name + name
}

type Car interface {
	honkTheHorn() string       // 摁喇叭
	getFlag() (string, string) // 获取车标
}

type Benz struct {
}
type Maserati struct {
	Color  string `form:"color"`
	Height int    `form:"height"`
}

func (benz *Benz) getFlag() (string, string) {
	return "三角标", "圆形"
}
func (benz *Benz) honkTheHorn() string {
	return "嘟嘟嘟"
}

func (ma *Maserati) getFlag() (string, string) {
	return "三叉戟", "异形"
}
func (ma *Maserati) honkTheHorn() string {
	return "嘀嘀嘀"
}

var sd int16

const (
	name = "sdsdsd"
)

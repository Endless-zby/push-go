package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// GlobalMiddleWare 和 Logger是两种不同的写法

func GlobalMiddleWare(context *gin.Context) {
	fmt.Println("中间件开始执行了")
	context.Set("all", "allwe")
	context.Next()
	fmt.Println("中间件执行结束了")
}

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

// 没有Next()函数的时候，该中间件会一次执行完然后跳转进入路由
func NoNextFunction(context *gin.Context) {
	fmt.Println("完全之后完后跳转路由的中间件--进入")
	fmt.Println("完全之后完后跳转路由的中间件--退出")
}

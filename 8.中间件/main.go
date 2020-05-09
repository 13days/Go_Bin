package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func checkLogin(ok bool)func(c *gin.Context)  {
	// 初始化工作...
	return func(c *gin.Context) {
		if ok{
			fmt.Println("正在验证登录....")
			// 验证成功放行
			// 是否登录的判断
			c.Next()
			// 验证失败拦截
			//c.Abort()
		}else{
			// 放行
			c.Next()
		}
	}
}

// 定义一个中间件m1:统计请求处理函数的耗时
func LoggerTime(c *gin.Context)  {
	fmt.Println("LoggerTime in...")
	// 计时
	start := time.Now()
	c.Set("user", gin.H{
		"user" : "admin",
		"password" : 123456,
	})
	// 异步执行
	//go funcXX(c.Copy())  // 在funcXX中只能使用c的拷贝
	c.Next()  // 调用后续的处理函数
	//c.Abort() // 阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("LoggerTime out...")
}

// 上下文数据获取
func getLoggerData(c *gin.Context)  {
	usr := c.MustGet("user").(gin.H)
	for k,v := range usr{
		fmt.Println(k,v)
	}
}
func main() {
	// 默认使用了Logger和Recovery中间件,日志,和recovery任何panic返回500
	//r := gin.Default()
	r := gin.New()

	// 全局中间件
	r.Use(checkLogin(true))

	// 单路由中间件
	r.GET("/index", LoggerTime, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg" : "getOk",
		})
	})

	// 多中间件注册
	r.GET("/more", LoggerTime, getLoggerData, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg" : "getOk",
		})
	})

	// 路由组中间件注册
	xxGroup := r.Group("/group", LoggerTime)
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
		})
	}
	r.Run()
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		// 方法1 : 使用map
		//data := map[string]interface{}{
		//	"name":"小明",
		//	"message":"hello world",
		//	"code":200,
		//}
		data := gin.H{
			"name":"小明",
			"message":"hello world",
			"code":200,
		}
		c.JSON(http.StatusOK, data)
	})

	// 方法2:结构体
	type msg struct{
		Name string `json:"name"`// 大写才能被前端访问,go序列化为Json,是用反射技术的,可以使用tag在做定制化操作
		Message string
		age int // 不可见
	}
	r.GET("/json_struct", func(c *gin.Context) {
		data := msg{
			"小明",
			"hello world",
			200,
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":9090")
}

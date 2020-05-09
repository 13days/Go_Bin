package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	// 加载静态文件
	r.Static("/xxx", "./static")

	// gin框架添加自定义函数safe
	r.SetFuncMap(template.FuncMap{
		"safe" : func(str string) template.HTML{
			return template.HTML(str)
		},
	})

	// 解析模板
	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")

	// 渲染模板
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	// 使用自定义函数
	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})


	r.Run(":8080")
}

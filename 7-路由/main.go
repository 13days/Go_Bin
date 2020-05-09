package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "get",
		})
	})
	
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "post",
		})
	})
	
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "put",
		})
	})
	
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "delete",
		})
	})
	
	// 使用any接收任意方法请求
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"message" : "get"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"message" : "post"})

			
		}
	})

	// 没有匹配到路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg" : "没有找到"})
	})

	// 路由组
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg" : "shop/index",
			})
		})
		shopGroup.GET("/cart", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg" : "shop/cart",
			})
		})
		shopGroup.POST("/checkout", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg" : "shop/checkout",
			})
		})
		// 路由组支持嵌套
		shopHeadGroup := shopGroup.Group("/head")
		{
			shopHeadGroup.GET("/get", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"msg" : "shop/head/get",
				})
			})
		}
	}
	r.Run()
}

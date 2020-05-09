package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Binding from JSON
type Login struct {
	// 大写为了能使用反射,tag使得查找能够一一对应起来
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// 获取浏览器参数,GET方法
	r.GET("/get", func(c *gin.Context) {
		// 查询
		str := c.Query("ping")
		// 查不到,使用默认值
		defaultStr := c.DefaultQuery("name", "somebody")
		// 带查询结果的
		qs, ok := c.GetQuery("age")
		data := gin.H{
			"pone" : str,
			"name" : defaultStr,
		}
		if ok {
			data["age"] = qs
		} else {
			data["age"] = "不知道多大"
		}
		c.JSON(http.StatusOK, data)
	})

	// 获取POST方法表单数据
	r.POST("/post", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		defaultName := c.DefaultPostForm("defaultName", "小王子")
		username := c.PostForm("username")
		address := c.PostForm("address")

		ps, ok := c.GetPostForm("message")
		if !ok {
			ps = "false"
		}
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"defaultName" : defaultName,
			"message":  ps,
			"username": username,
			"address":  address,
		})
	})

	// 获取uri路径参数
	// 请求的参数通过URL路径传递，例如：/path/小王子/沙河
	// /:username/:address 不要使用这种,会路由匹配冲突
	r.GET("/path/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")

		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	// path常用场景
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")

		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"year": year,
			"month":  month,
		})
	})

	// 参数绑定,不用自己每次去解析URL的数据
	//ShouldBind会按照下面的顺序解析请求中的数据完成绑定：
	//
	//如果是 GET 请求，只使用 Form 绑定引擎（query）。
	//如果是 POST 请求，首先检查 content-type 是否为 JSON 或 XML，然后再使用 Form（form-data）。

	// 绑定JSON的示例 ({"user": "q1mi", "password": "12346"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 编码和JSON差不多...
	// 绑定form表单示例 (user=q1mi&password=123456)
	r.POST("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定QueryString示例 (/loginQuery?user=q1mi&password=123456)
	// localhost:9090/loginForm?user=q1mi&password=123456
	r.GET("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	r.Run(":9090")
}

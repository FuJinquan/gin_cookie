package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//授权中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("abc")
		if err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		//多与授权中间件连用，若授权验证失败，则Abort以确保不调用此请求接下来的的处理程序。
		c.Abort()
		return
	}
}

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		//key_cookie string, cookie名
		//value_cookie string, cookie值
		// maxAge int, 单位为秒
		// path,cookie所在目录
		// domain string,域名
		// secure 是否智能通过https访问
		// httpOnly bool  是否允许别人通过js获取自己的cookie
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		c.String(200, "Login Successful!")
	})
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
	r.Run(":8080")
}

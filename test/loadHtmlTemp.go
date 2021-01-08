package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadHtmlTemplate() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "Main website",
			"message": "瓦蓝蓝的天上飞🦅，我在东北眺望北京",
		})
	})
	router.Run(":8080")
}

func LoadHtmlDiffPath() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
			"message": "瓦蓝蓝的天上飞🦅，我在东北眺望北京",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
			"message": "瓦蓝蓝的天上飞🦅，我在东北眺望北京",
		})
	})
	router.Run(":8080")

}

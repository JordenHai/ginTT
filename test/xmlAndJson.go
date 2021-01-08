package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func XmlAndJsonTest() *gin.Engine{
	r := gin.Default()

	r.GET("/someJson", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"honey",
			"status":http.StatusOK,
		})
	})

	r.GET("/moreJson", func(c *gin.Context) {
		var msg struct{
			User     string `json:"user"`
			Message  string `json:"message"`
			Number   int    `json:"number"`
		}

		msg.Number = 12345
		msg.User = "long"
		msg.Message ="瓦蓝蓝的天上飞老🦅"

		c.JSON(http.StatusOK,msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"我在东北眺望北京",
			"status":http.StatusOK,
		})
	})

	return r
}
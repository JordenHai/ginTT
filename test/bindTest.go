package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`

}

func bindTests(){
	route := gin.Default()


	route.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if c.BindJSON(&json) == nil{
			if json.Username == "jiaohai" && json.Password == "123456"{
				c.JSON(http.StatusOK,gin.H{
					"status":"Login success!",
				})
			}else {
				c.JSON(http.StatusBadRequest,gin.H{
					"status":"Login failed",
				})
			}
		}
	})


	route.POST("/loginForm", func(c *gin.Context) {
		var json Login
		if c.Bind(&json) == nil{
			if json.Username == "jiaohai" && json.Password == "123456"{
				c.JSON(http.StatusOK,gin.H{
					"status":"Login success!",
				})
			}else {
				c.JSON(http.StatusBadRequest,gin.H{
					"status":"Login failed",
				})
			}
		}
	})

	route.Run()
}
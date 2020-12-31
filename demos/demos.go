package demos

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//var context *gin.Context

func loginEndpoint(c *gin.Context){
	c.String(200,"login")
	fmt.Println("这是登陆入口")
}

func submitEndpoint(c *gin.Context){
	fmt.Println("这是提交入口")
}

func readEndpoint(c *gin.Context){
	fmt.Println("这是读取入口")
}

func routeTest(){
	route := gin.Default()
	v1:=route.Group("/v1")
	{
		v1.GET("/login",loginEndpoint)
		v1.GET("/submit",submitEndpoint)
		v1.GET("/readEndpoint",readEndpoint)
	}

	v2:=route.Group("/v2")
	{
		v2.GET("/login",loginEndpoint)
		v2.GET("/submit",submitEndpoint)
		v2.GET("/readEndpoint",readEndpoint)
	}

	route.Run()

}

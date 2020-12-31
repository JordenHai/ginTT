package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func multiUpload(){
	route := gin.Default()
	route.POST("/uploads", func(context *gin.Context) {

		if form,err := context.MultipartForm();err != nil{
			context.String(http.StatusBadRequest,"Error")
			return
		}else {
			files := form.File["upload[]"]
			for id,file := range files{
				if errs := context.SaveUploadedFile(file,file.Filename);errs != nil{
					context.String(http.StatusBadRequest,"the process of the %d has been failed",id)
				}else {
					context.String(http.StatusOK,"the process of %d/%d\n",id,len(files))
				}
			}

		}
	})
	route.Run()
}

func upload(){
	route := gin.Default()
	route.POST("/upload", func(context *gin.Context) {
		if file,err := context.FormFile("upload");err != nil{
			context.String(http.StatusBadRequest,"Request Failed")
			return
		}else {
			filename := file.Filename
			fmt.Println("Filename",filename)

			if errs := context.SaveUploadedFile(file,filename);errs != nil{
				context.String(http.StatusBadRequest, "保存失败 Error:%s", errs.Error())
				return
			}else {
				context.String(http.StatusOK,"Upload Success")
			}
		}
	})
	route.Run()
}

func withGetAndPost(){
	route := gin.Default()
	route.POST("/getAndpost", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page","0")
		name := context.PostForm("name")
		message := context.PostForm("message")
		fmt.Printf("Get方式:id: %s; page: %s; Post方式:name: %s; message: %s", id, page, name, message)
	})

	route.Run()
}

func withValueRoute(){
	route := gin.Default()
	
	route.GET("/wel", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname","Guest")
		lastname := context.Query("lastname")
		context.String(http.StatusOK, "hello %s %s",firstname,lastname)
	})

	route.POST("/from_post", func(context *gin.Context) {
		msg:=context.PostForm("message")
		nick := context.DefaultPostForm("nick","anonymous")
		context.JSON(http.StatusOK,gin.H{
			"status":"ok",
			"message":msg,
			"nick":nick,
		})
	})
	route.Run()
}

func simpleRoute(){
	route := gin.Default()

	route.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK,"your name "+name)
	})

	route.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name +" "+action
		context.String(http.StatusOK,message)

	})
	route.Run()
}

func test(){

	route := gin.Default()

	route.GET("/",func(ctx *gin.Context){
		ctx.String(http.StatusOK,"hello world")
	})

	route.Run()
}

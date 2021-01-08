package test

import "net/http"
import "github.com/gin-gonic/gin"

func StaticDock() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Listen and server on 0.0.0.0:8080
	router.Run(":8080")
}


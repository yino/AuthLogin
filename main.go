package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yino/AuthLogin"
)

func main() {
	router := gin.Default()
	router.GET("/gitee/getToken", AuthLogin.GiteenGetToken)
	router.GET("/gitee/login", AuthLogin.GiteeLogin)

	router.Run(":9527")
}

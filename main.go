package main

import (
	"AuthLogin/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/gitee/getToken", auth.GiteenGetToken)
	router.GET("/gitee/login", auth.GiteeLogin)
	router.GET("/gitee/tokenRollback", auth.GiteeRollback)
	router.Run("localhost:9527")
}

package	github.com/yino/AuthLogin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GiteenGetToken(c *gin.Context){
	clientId := "a28944a66fd8a0928a4ce5224ff5289c2cdcfa25a4f5057cccb99f561ce6b89b"
	redirectUrl := "http://localhost:9527/gitee/login"
	authUrl := fmt.Sprintf("https://gitee.com/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code", clientId, redirectUrl)
	c.Redirect(http.StatusMovedPermanently, authUrl)
}
func GiteeLogin(c *gin.Context)  {
		
}
package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	r "github.com/solos/requests"
)

var clientId = "a28944a66fd8a0928a4ce5224ff5289c2cdcfa25a4f5057cccb99f561ce6b89b"
var clientSecret = "f2c39ed9ccbcc85018c23c612e77d6f28e9ffc5ecc8de42fbe66eb32575e675f"
var getCodeRollback = "http://localhost:9527/gitee/login"
var getTokenRollback = "http://localhost:9527/gitee/tokenRollback"

func GiteenGetToken(c *gin.Context) {

	authUrl := fmt.Sprintf("https://gitee.com/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=user_info%20projects%20pull_requests", clientId, getCodeRollback)
	c.Redirect(http.StatusMovedPermanently, authUrl)
}

func GiteeLogin(c *gin.Context) {
	code := c.Query("code")
	accessToeknUrl := fmt.Sprintf("https://gitee.com/oauth/token?grant_type=authorization_code&code=%s&client_id=%s&redirect_uri=%s&client_secret=%s", clientId, getTokenRollback)

	kwargs := r.M{}
	kwargs["timeout"] = 10

	// 设定POST请求数据
	data := make(map[string]string)
	req := &r.Request{Args: kwargs}
	resp, _ := req.MakeRequest("POST", accessToeknUrl, r.Headers(data), r.Data(data))
	fmt.Println("resp:", resp.Content)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": code,
	})
}

func GiteeRollback(c *gin.Context) {

}

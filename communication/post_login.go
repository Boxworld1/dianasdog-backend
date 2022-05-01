// @title	PostLogin
// @description	后端密码接口
// @auth	ryl		2022/4/21	10:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/database"

	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	Username string `json:"username" binding:"required"`
}

// PostLogin postlogin
// @Summary 后端密码接口
// @Description 后端密码接口
// @Accept json
// @Produce json
// @Param username query string true "用户名"
// @Success 200 {object} string "{"msg": "hello Razeen"}"
// @Failure 400 {object} string "{"msg": "who are you"}"
// @Router /login [post]
func PostLogin(context *gin.Context) {
	var body LoginBody

	// 检查收到信息的格式是否正确
	err := context.ShouldBindJSON(&body)

	// 若不是，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 取得 username 字段
	username := body.Username
	pwd, err := database.UserSignIn(username)

	// 查找失败
	if pwd == "None" || err != nil {
		context.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 返回结果
	context.JSON(200, gin.H{
		"password": pwd,
	})
}

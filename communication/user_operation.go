// @title	OperateUser
// @description	后端添加和删除用户接口
// @auth	王梓桥		2022/5/9	17:18
// @param	context	*gin.Context

package communication

import (
	"dianasdog/database"

	"github.com/gin-gonic/gin"
)

type DeleteBody struct {
	Username string `json:"username" binding:"required"`
}

// @Summary 删除用户接口
// @Tags User
// @Description 删除用户
// @Accept json
// @Produce json
// @Param username query string true "用户名"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad Request"
// @Router /userdelete [post]
func DeleteUser(context *gin.Context) {
	var body DeleteBody

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
	err = database.DeleteUser(username)

	// 删除失败
	if err != nil {
		context.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 返回结果
	context.JSON(200, gin.H{
		"response": "succeed",
	})
}

type AddBody struct {
	Username     string `json:"username" binding:"required"`
	UserPassword string `json:"userpassword" binding:"required"`
	UserLevel    string `json:"userlevel" binding:"required"`
}

// @Summary 增加用户接口
// @Tags User
// @Description 增加用户
// @Accept json
// @Produce json
// @Param username query string true "用户名"
// @Param userpassword query string true "密码"
// @Param userlevel query string true "用户等级"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad Request"
// @Router /useradd [post]
func AddUser(context *gin.Context) {
	var body AddBody
	// 检查收到信息的格式是否正确
	err := context.ShouldBindJSON(&body)
	// 若不是，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 取得 user相关信息
	username := body.Username
	userpassword := body.UserPassword
	userlevel := body.UserLevel
	err = database.InsertPwdIntoSQL(userpassword, username, userlevel)

	// 添加失败
	if err != nil {
		context.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 返回结果
	context.JSON(200, gin.H{
		"response": "succeed",
	})
}

// @Summary 取得所有用户名
// @Tags User
// @Description 取得所有用户名
// @Produce json
// @Success 200 {object} string "{"data": ["name1", "name2"]}"
// @Failure 400 {object} string "Bad Request"
// @Router /alluser [get]
func AllUser(context *gin.Context) {
	data, err := database.AllUser()

	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"data": data,
	})
}

// @title	GetTemplate
// @description	后端发出配置之接口
// @auth	ryl		2022/4/26	23:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/database"

	"github.com/gin-gonic/gin"
)

type GetTemplateBody struct {
	Resource string `json:"resource" binding:"required"`
}

func GetTemplate(context *gin.Context) {
	var body GetTemplateBody

	// 检查收到信息的格式是否正确
	err := context.ShouldBindJSON(&body)

	// 若不是，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 取得文件
	data, err := database.GetFile(database.TemplateClient, "file", body.Resource)

	// 若不存在文件/对应特型卡类型，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 否则正常返回结果
	context.JSON(200, gin.H{
		"data": string(data),
	})
}

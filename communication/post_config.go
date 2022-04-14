// @title	PostConfig
// @description	后端接收配置文件之接口
// @auth	ryl		2022/4/14	10:30
// @param	context	*gin.Context

package communication

import (
	"github.com/gin-gonic/gin"
)

type ConfigBody struct {
	Resource string                 `json:"resource" binding:"required"`
	Data     map[string]interface{} `json:"data" binding:"required"`
}

func PostConfig(context *gin.Context) {
	var body ConfigBody

	// 检查收到信息的格式是否正确
	err := context.ShouldBindJSON(&body)

	// 若不是，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 取得特型卡类型及其内容
	res := body.Resource

	// 返回对应信息
	context.JSON(200, gin.H{
		"content": res, //result,
	})
}

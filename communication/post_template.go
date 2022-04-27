// @title	PostTemplate
// @description	后端接收配置文件之接口
// @auth	ryl		2022/4/20	18:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/setter"

	"github.com/gin-gonic/gin"
)

type TemplateBody struct {
	Resource  string `form:"resource" binding:"required"`
	Pattern   string `form:"pattern" binding:"required"`
	Operation string `form:"operation" binding:"required"`
}

func PostTemplate(context *gin.Context) {
	var body TemplateBody

	// 检查收到信息的格式是否正确
	err := context.ShouldBindJSON(&body)

	// 若不是，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 取得特型卡 ID
	res := body.Resource
	pat := body.Pattern
	opt := body.Operation

	// 否则调用函数写入文件
	setter.SetTemplate(res, pat, opt)

	// 返回对应值
	context.JSON(200, gin.H{
		"message": "successful!", //result,
	})
}

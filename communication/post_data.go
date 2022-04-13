// @title	PostData
// @description	后端接收数据用接口
// @auth	ryl		2022/4/13		17:30
// @param	context	*gin.Context

package communication

import (
	"github.com/gin-gonic/gin"
)

type DataBody struct {
	Type     string                 `json:"type" binding:"required"`
	Resource string                 `json:"resource" binding:"required"`
	File     string                 `json:"file" binding:"required"`
	Data     map[string]interface{} `json:"data" binding:"required"`
}

func PostData(context *gin.Context) {
	var body DataBody
	err := context.ShouldBindJSON(&body)
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	typ := body.Type
	context.JSON(200, gin.H{
		"content": typ, //result,
	})
}

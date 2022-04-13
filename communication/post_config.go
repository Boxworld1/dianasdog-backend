// @title	PostConfig
// @description	后端接收配置文件之接口
// @auth	ryl		2022/4/13	13:30
// @param	context	*gin.Context

package communication

import (
	"github.com/gin-gonic/gin"
)

type ConfigBody struct {
	Type string `json:"type" binding:"required"`
}

func PostConfig(context *gin.Context) {
	var body ConfigBody
	err := context.ShouldBindJSON(&body)
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	query := body.Type
	// result := search.IntentRecognition(query)
	context.JSON(200, gin.H{
		"content": query, //result,
	})
}

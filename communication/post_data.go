// @title	GoSearch
// @description	后端接收数据用接口
// @auth	ryl		2022/4/13		13:30
// @param	context	*gin.Context

package communication

import (
	"github.com/gin-gonic/gin"
)

type DataBody struct {
	Type string `json:"type" binding:"required"`
	Word string `json:"word" binding:"required"`
	File string `json:"file" binding:"required"`
	Data string `json:"data" binding:"required"`
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

	query := body.Type
	// result := search.IntentRecognition(query)
	context.JSON(200, gin.H{
		"content": query, //result,
	})
}

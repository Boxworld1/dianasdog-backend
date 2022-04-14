// @title	GoSearch
// @description	后端搜索接口
// @auth	ryl		2022/4/14	10:30
// @param	context	*gin.Context

package communication

import (
	"github.com/gin-gonic/gin"
)

type SearchBody struct {
	Query string `json:"query" binding:"required"`
}

func GoSearch(context *gin.Context) {
	var body SearchBody

	// 检查收到信息的格式是否正确
	err := context.ShouldBindJSON(&body)

	// 若不是，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 取得 query 字段
	query := body.Query

	// 开始搜索流程
	// result := search.IntentRecognition(query)

	// 返回结果
	context.JSON(200, gin.H{
		"content": query, //result,
	})
}

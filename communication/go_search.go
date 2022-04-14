// @title	GoSearch
// @description	后端搜索接口
// @auth	ryl		2022/4/13	13:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/setup"
	"fmt"

	"github.com/gin-gonic/gin"
)

type SearchBody struct {
	Query string `json:"query" binding:"required"`
}

func GoSearch(context *gin.Context) {
	var body SearchBody
	fmt.Println(body)
	err := context.ShouldBindJSON(&body)
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	query := body.Query
	fmt.Println(query)
	query, _ = setup.GetAbsPath()
	// result := search.IntentRecognition(query)
	context.JSON(200, gin.H{
		"content": query, //result,
	})
}

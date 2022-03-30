// @title		Router
// @description	后端与前端交互之接口
// @auth		ryl				2022/3/30		16:30

package communication

import (
	"write_setting/search_service"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Query string `json:"query" binding:"required"`
}

func Router() {
	router := gin.Default()
	router.GET("", func(context *gin.Context) {
		var body Body
		err := context.ShouldBindJSON(&body)
		if err != nil {
			context.JSON(400, gin.H{
				"err": err.Error(),
			})
			return
		}

		query := body.Query
		result := search_service.SearchFunction(query)
		context.JSON(200, gin.H{
			"content": result,
		})
	})
	router.Run()
}

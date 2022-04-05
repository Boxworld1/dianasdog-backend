// @title		SetupRouter
// @description	后端与前端交互之接口
// @auth		ryl				2022/3/30		22:30

package communication

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Query string `json:"query" binding:"required"`
}

func GoSearch(context *gin.Context) {
	var body Body
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
	// result := search_service.SearchFunction(query)
	context.JSON(200, gin.H{
		// "content": result,
		"content": query,
	})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/search", GoSearch)
	// router.Run()
	return router
}

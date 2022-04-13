// @title		SetupRouter
// @description	后端与前端交互之接口
// @auth		ryl				2022/4/6		23:30

package communication

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/search", GoSearch)
	router.POST("/data", PostData)
	router.POST("/pattern", PostConfig)
	router.POST("/setting", PostSetting)

	return router
}

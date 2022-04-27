// @title	GetCategory
// @description	后端回传所有特型卡名字
// @auth	ryl		2022/4/27		21:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/database"

	"github.com/gin-gonic/gin"
)

func GetCategory(context *gin.Context) {

	// 取得文件
	data, _ := database.GetAllCategory(database.CategoryClient, "word")

	// 返回结果
	context.JSON(200, gin.H{
		"data": data,
	})
}

// @title	GetCategory
// @description	后端回传所有特型卡名字
// @auth	ryl		2022/4/27		21:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/database"

	"github.com/gin-gonic/gin"
)

// @Summary 取得现有特型卡类型
// @Tags Setting
// @Description 后端返回现有特型卡类型之接口
// @Produce json
// @Success 200 {object} string "{"data": data}"
// @Failure 400 {object} string "Bad Request"
// @Router /category [get]
func GetCategory(context *gin.Context) {

	// 取得文件
	data, _ := database.GetAllCategory(database.CategoryClient, "word")

	// 返回结果
	context.JSON(200, gin.H{
		"data": data,
	})
}

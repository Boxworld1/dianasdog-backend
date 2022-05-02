// @title	GetKey
// @description	取得某个特型卡对应 item 中的所有键值
// @auth	ryl		2022/4/26		17:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/database"

	"github.com/gin-gonic/gin"
)

// @Summary 取得数据键值
// @Description 后端返回某一数据类型键值之接口
// @Produce json
// @Param resource query string true "特型卡名称 (如: car, poem 等)"
// @Success 200 {object} []string "{"data": data}"
// @Failure 400 {object} string "Bad Request"
// @Router /key [get]
func GetKey(context *gin.Context) {

	// 检查收到信息的格式是否正确
	resource, ok := context.GetQuery("resource")

	// 若不是，则返回错误
	if !ok {
		context.JSON(400, gin.H{
			"err": "wrong param",
		})
		return
	}

	// 取得文件
	data, err := database.GetAllCategory(database.CategoryClient, resource)

	// 若不存在文件/对应特型卡类型，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 否则正常返回结果
	context.JSON(200, gin.H{
		"data": data,
	})
}

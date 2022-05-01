// @title	GetDataName
// @description	后端发出写入行为之接口
// @auth	ryl		2022/4/26		19:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/database"

	"github.com/gin-gonic/gin"
)

type GetDataNameBody struct {
	Resource string `json:"resource" binding:"required"`
}

// @Summary 取得数据文件名
// @Description 后端返回数据文件名之接口
// @Accept json
// @Produce json
// @Param resource query string true "特型卡名称 (如: car, poem 等)"
// @Success 200 {object} []string "文件名"
// @Failure 400 {object} string "Bad Request"
// @Router /dataname [get]
func GetDataName(context *gin.Context) {
	var body GetDataNameBody

	// 检查收到信息的格式是否正确
	err := context.ShouldBindJSON(&body)

	// 若不是，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 取得文件
	data, err := database.GetFileName(database.DataClient, body.Resource)

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

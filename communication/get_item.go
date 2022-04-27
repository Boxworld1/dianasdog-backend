// @title	GetItem
// @description	后端返回 item 用接口
// @auth	ryl		2022/4/27		10:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/database"

	"github.com/gin-gonic/gin"
)

type GetItemBody struct {
	Resource string `json:"resource" binding:"required"`
	Key      string `json:"key" binding:"required"`
}

func GetItem(context *gin.Context) {
	var body GetItemBody

	// 检查收到信息的格式是否正确
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 若无错误，则继续
	resource := body.Resource
	key := body.Key
	docid := resource + "@" + key

	data, err := database.GetFile(database.DocidClient, resource, docid)

	// 不存在此特型卡类型或文件
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"data": string(data),
	})

}

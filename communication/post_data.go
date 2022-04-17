// @title	PostData
// @description	后端接收数据用接口
// @auth	ryl		2022/4/13		17:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/io"

	"github.com/gin-gonic/gin"
)

type DataBody struct {
	Type     string `json:"type" binding:"required"`
	Resource string `json:"resource" binding:"required"`
	File     string `json:"file" binding:"required"`
	Data     string `json:"data" binding:"required"`
}

func PostData(context *gin.Context) {
	var body DataBody

	// 检查收到信息的格式是否正确
	err := context.ShouldBindJSON(&body)

	// 若不是，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	typ := body.Type
	res := body.Resource
	filename := body.File
	content := body.Data

	// 否则按照操作类型进行操作
	switch typ {
	// 写入文件
	case "insert":
		// 将内容转化为 []byte 方便写入文件
		io.SetData(res, filename, []byte(content))
	// 删除条目
	case "delete":

	// 更新条目
	case "update":

	}

	// 返回对应值
	context.JSON(200, gin.H{
		"message": "successful!", //result,
		"data":    content,
	})
}

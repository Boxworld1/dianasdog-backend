// @title	GetData
// @description	后端接收数据用接口
// @auth	ryl		2022/4/20		0:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetDataBody struct {
	Resource string `json:"resource" binding:"required"`
	Filename string `json:"filename" binding:"required"`
}

func GetData(context *gin.Context) {
	var body GetDataBody

	// 检查收到信息的格式是否正确
	if err := context.ShouldBind(&body); err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 若无错误，则继续
	resource := body.Resource
	filename := body.Filename

	data, err := database.GetFile(database.DataClient, resource, filename)

	// 不存在此特型卡类型
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	contentType := "text/xml"
	fileContentDisposition := "attachment;filename=\"" + filename + "\""
	context.Status(200)
	context.Header("Content-Type", contentType)
	context.Header("Content-Disposition", fileContentDisposition)
	context.Data(http.StatusOK, contentType, data)

}

// @title	GetData
// @description	后端返回数据用接口
// @auth	ryl		2022/4/26		0:30
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

// @Summary 取得数据
// @Description 后端返回数据之接口
// @Produce mpfd
// @Param resource query string true "特型卡名称 (如: car, poem 等)"
// @Param filename query string true "文件名"
// @Success 200 {file} string "XML 文件"
// @Failure 400 {object} string "Bad Request"
// @Router /data [get]
func GetData(context *gin.Context) {

	// 检查收到信息的格式是否正确
	resource, ok1 := context.GetQuery("resource")
	filename, ok2 := context.GetQuery("filename")

	// 若不是，则返回错误
	if !ok1 || !ok2 {
		context.JSON(400, gin.H{
			"err": "wrong param",
		})
		return
	}

	// 若无错误，则继续
	data, err := database.GetFile(database.DataClient, resource, filename)

	// 不存在此特型卡类型或文件
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

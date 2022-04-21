// @title	PostData
// @description	后端接收数据用接口
// @auth	ryl		2022/4/20		0:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/database"
	"errors"
	"io/ioutil"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type DataBody struct {
	Type     string                `form:"type" binding:"required"`
	Resource string                `form:"resource" binding:"required"`
	Filename string                `form:"filename"`
	File     *multipart.FileHeader `form:"file"`
	Data     string                `form:"data"`
}

func PostData(context *gin.Context) {
	var body DataBody
	var err error
	var msg string

	// 检查收到信息的格式是否正确
	err = context.ShouldBind(&body)

	// 若不是，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 若无错误，则继续
	typ := body.Type
	res := body.Resource
	var filename string
	var data []byte

	// 取得对应的数据内容
	if len(body.Data) > 0 {
		// 若是字符串
		content := body.Data
		filename = body.Filename
		data = []byte(content)

		if len(filename) <= 0 {
			msg = "invaild filename!"
		}

	} else if body.File != nil {
		// 若使用文件传输
		filename = body.File.Filename
		fileContent, _ := body.File.Open()
		data, err = ioutil.ReadAll(fileContent)

		if err != nil {
			msg = err.Error()
		}

	} else {
		// 若没有传输数据，则错误
		msg = "form data error: wrong parameters!"
	}

	// 若过程中出现错误
	if len(msg) > 0 {
		context.JSON(400, gin.H{
			"err": msg,
		})
		return
	}

	err = errors.New("form data error")
	// 否则按照操作类型进行操作
	switch typ {
	// 写入文件
	case "insert":
		// 将内容转化为 []byte 方便写入文件
		// err = io.SetData(res, filename, data)
		database.CreateFileTable(database.DataClient, res)
		database.InsertFile(database.DataClient, res, filename, data)
		// setup.UnpackXmlFile(filename, res)
		err = nil
	// 删除条目
	case "delete":
		err = nil
	// 更新条目
	case "update":
		err = nil
	}

	// 返回对应值
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
	} else {
		context.JSON(200, gin.H{
			"message": "successful!", //result,
		})
	}

}

// @title	PostConfig
// @description	后端接收配置文件之接口
// @auth	ryl		2022/4/14	10:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/io"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ConfigBody struct {
	Resource string                `form:"resource" binding:"required"`
	Data     string                `form:"data" binding:"-"`
	File     *multipart.FileHeader `form:"file" binding:"-"`
}

type ConfigJson struct {
	Resource string                 `form:"resource" binding:"required"`
	Data     map[string]interface{} `form:"data" binding:"required"`
}

func PostConfig(context *gin.Context) {
	var body ConfigBody
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

	// 取得特型卡 ID
	res := body.Resource
	var data []byte

	// 取得对应的数据内容
	if len(body.Data) > 0 {
		// 若使用 json 格式的数据
		content := body.Data

		// 去掉前端多余的引号
		str, _ := strconv.Unquote(content)

		// 检查数据内容是否正确
		var jsonContent ConfigJson
		err = json.Unmarshal([]byte(str), &jsonContent)

		// 若不正确，则返回错误
		if err != nil {
			msg = err.Error()
		}

		if jsonContent.Data == nil {
			msg = "json data error: wrong parameters!"
		}

		data, _ = json.Marshal(jsonContent)

	} else if body.File != nil {
		// 若使用文件传输
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

	// 否则调用函数写入文件
	io.SetTemplate(res, data)

	// 返回对应值
	context.JSON(200, gin.H{
		"message": "successful!", //result,
	})
}

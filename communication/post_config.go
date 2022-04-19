// @title	PostConfig
// @description	后端接收配置文件之接口
// @auth	ryl		2022/4/14	10:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/io"
	"encoding/json"
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type ConfigBody struct {
	Resource string                `form:"resource" binding:"required"`
	Data     string                `form:"data" binding:"-"`
	File     *multipart.FileHeader `form:"file" binding:"-"`
}

type ConfigJson struct {
	Resource string                 `form:"resource" binding:"required"`
	Data     map[string]interface{} `form:"write_setting" binding:"required"`
}

func PostConfig(context *gin.Context) {
	var body ConfigBody

	// 检查收到信息的格式是否正确
	err := context.ShouldBind(&body)

	// 若不是，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 取得特型卡类型及其内容
	res := body.Resource
	content := body.Data

	var jsonContent ConfigJson
	err = json.Unmarshal([]byte(content), &jsonContent)
	fmt.Println(jsonContent)

	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 调用函数写入文件
	io.SetTemplate(res, []byte(content))

	// 返回对应信息
	context.JSON(200, gin.H{
		"message": "successful!", //result,
	})
}

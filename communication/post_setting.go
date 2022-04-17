// @title	GoSetting
// @description	后端接收写入行为之接口
// @auth	ryl		2022/4/13		17:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/io"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type SettingBody struct {
	Resource string                 `json:"resource" binding:"required"`
	Setting  map[string]interface{} `json:"write_setting" binding:"required"`
}

func PostSetting(context *gin.Context) {
	var body SettingBody

	// 检查收到信息的格式是否正确
	err := context.ShouldBindJSON(&body)

	// 若不是，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 取得特型卡 ID 及对应内容
	res := body.Resource
	content := body.Setting

	// 将内容转化为 []byte 方便写入文件
	str, _ := json.Marshal(content)

	// 调用函数写入文件
	io.SetConfig(res, str)

	// 返回对应值
	context.JSON(200, gin.H{
		"message": "successful!", //result,
	})
}

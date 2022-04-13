// @title	GoSetting
// @description	后端接收写入行为之接口
// @auth	ryl		2022/4/13		17:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/setup"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type SettingBody struct {
	Resource string                 `json:"resource" binding:"required"`
	Setting  map[string]interface{} `json:"write_setting" binding:"required"`
}

func PostSetting(context *gin.Context) {
	var body SettingBody
	err := context.ShouldBindJSON(&body)
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	res := body.Resource
	content := body.Setting

	str, err := json.Marshal(content)

	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	setup.SetConfig(res, str)

	context.JSON(200, gin.H{
		"message": "successful!", //result,
	})
}

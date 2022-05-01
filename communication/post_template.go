// @title	PostTemplate
// @description	后端接收配置文件之接口
// @auth	ryl		2022/4/20	18:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/setter"
	"errors"

	"github.com/gin-gonic/gin"
)

type TemplateBody struct {
	Type      string   `json:"type" binding:"required"`
	Resource  string   `json:"resource" binding:"required"`
	Data      []string `json:"data" binding:"required"`
	Operation string   `json:"operation" binding:"required"`
}

// @Summary 接收配置文件
// @Description 后端接收配置文件之接口
// @Accept mpfd
// @Produce json
// @Param type query string true "上传信息类型 (intent, garbage, pattern)"
// @Param resource query string true "特型卡名称 (如: car, poem 等)"
// @Param data query []string true "数据内容"
// @Param operation query string true "操作类型 (insert, delete)"
// @Success 200 {object} string "{"message": "successful!"}"
// @Failure 400 {object} string "错误原因"
// @Router /pattern [post]
func PostTemplate(context *gin.Context) {
	var body TemplateBody

	// 检查收到信息的格式是否正确
	err := context.ShouldBindJSON(&body)

	// 若不是，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 取得特型卡 ID, 操作类型 和 数据
	res := body.Resource
	data := body.Data
	opType := body.Operation
	wordType := body.Type

	// 调用函数写入文件
	switch wordType {
	case "pattern":
		setter.SetTemplate(res, data, opType)
	case "intent":
		setter.SetWord(res, data, opType, wordType)
	case "garbage":
		setter.SetWord(res, data, opType, wordType)
	default:
		err = errors.New("wrong word type")
	}

	if err != nil {
		context.JSON(400, gin.H{
			"err": err,
		})
		return
	}

	// 返回对应值
	context.JSON(200, gin.H{
		"message": "successful!", //result,
	})
}

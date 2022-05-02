// @title	GetTemplate
// @description	后端发出配置之接口
// @auth	ryl		2022/4/26	23:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/database"

	"github.com/gin-gonic/gin"
)

// @Summary 取得配置文件
// @Description 后端返回配置文件之接口
// @Produce json
// @Param resource query string true "特型卡名称 (如: car, poem 等)"
// @Param type query string true "下载信息类型 (intent, garbage, pattern)"
// @Success 200 {object} []string "{"data": data}"
// @Failure 400 {object} string "Bad Request"
// @Router /pattern [get]
func GetTemplate(context *gin.Context) {

	// 检查收到信息的格式是否正确
	res, ok1 := context.GetQuery("resource")
	wordType, ok2 := context.GetQuery("type")

	// 若不是，则返回错误
	if !ok1 || !ok2 {
		context.JSON(400, gin.H{
			"err": "wrong param",
		})
		return
	}

	var data []string
	var err error

	switch wordType {
	case "pattern":
		data, err = database.FetchAllPattern(res)
	default:
		data, err = database.GetAllWordFromDict(res, wordType)
	}

	// 若不存在文件/对应特型卡类型，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 否则正常返回结果
	context.JSON(200, gin.H{
		"data": data,
	})
}

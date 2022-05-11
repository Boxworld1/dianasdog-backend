// @title	GetItem
// @description	后端返回 item 用接口
// @auth	ryl		2022/4/27		10:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/database"

	"github.com/beevik/etree"
	"github.com/gin-gonic/gin"
)

// @Summary 取得单一数据
// @Tags Data
// @Description 后端返回某一条数据之接口
// @Produce json
// @Param resource query string true "特型卡名称 (如: car, poem 等)"
// @Param key query string true "索引值 (默认为 item.key)"
// @Success 200 {object} string "{"data": data}"
// @Failure 400 {object} string "Bad Request"
// @Router /item [get]
func GetItem(context *gin.Context) {

	// 检查收到信息的格式是否正确
	resource, ok1 := context.GetQuery("resource")
	key, ok2 := context.GetQuery("key")

	// 若不是，则返回错误
	if !ok1 || !ok2 {
		context.JSON(400, gin.H{
			"err": "wrong param",
		})
		return
	}

	// 若无错误，则继续
	docid := resource + "@" + key

	data, err := database.GetDocid(database.DocidClient, resource, docid)

	// 不存在此特型卡类型或文件
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 数据转化为 xml
	doc := etree.NewDocument()
	doc.ReadFromBytes(data)
	str, _ := doc.WriteToString()
	str = "<DOCUMENT>\n" + str + "\n</DOCUMENT>"

	// 正常返回结果
	context.JSON(200, gin.H{
		"data": str,
	})

}

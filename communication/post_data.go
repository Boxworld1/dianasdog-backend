// @title	PostData
// @description	后端接收数据用接口
// @auth	ryl		2022/4/20		0:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/database"
	"dianasdog/setter"
	"dianasdog/setup"
	"io/ioutil"
	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"
)

type DataBody struct {
	Type     string                `form:"type" binding:"required"`
	Resource string                `form:"resource" binding:"required"`
	Filename string                `form:"filename"`
	File     *multipart.FileHeader `form:"file"`
	Data     string                `form:"data"`
	Key      string                `form:"key"`
}

// @Summary 发送 XML 数据
// @Tags Data
// @Description 后端接收 XML 数据之接口。若为插入，则需要有 data 或 file 中的一个；若为删除，则需要 key 或 filename 中的一个。（若同时出现则两者都删除）
// @Accept mpfd
// @Produce json
// @Param resource formData string true "特型卡名称 (如: car, poem 等)"
// @Param type formData string true "操作类型 (insert, delete)"
// @Param data formData string false "数据内容"
// @Param filename formData string false "文件名"
// @Param file formData file false "文件形式上传之数据"
// @Param key formData string false "要删除的 key (即不含 resource 的 docid, 如: mytest@test 中的 test)"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad Request"
// @Router /data [post]
func PostData(context *gin.Context) {
	var body DataBody
	var err error
	var msg string

	// 检查收到信息的格式是否正确
	if err = context.ShouldBind(&body); err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 若无错误，则继续
	typ := body.Type
	res := body.Resource
	filename := body.Filename
	key := body.Key
	var data []byte

	// 取得对应的数据内容
	if len(body.Data) > 0 {
		// 若是字符串
		content := body.Data
		content = strings.Replace(content, "\\n", "\n", -1)
		content = strings.Replace(content, "\\t", "\t", -1)
		content = strings.Replace(content, "\\r", "\r", -1)
		data = []byte(content)

		if len(filename) <= 0 {
			filename = "."
		}

	} else if body.File != nil {
		// 若使用文件传输
		filename = body.File.Filename
		fileContent, _ := body.File.Open()
		if data, err = ioutil.ReadAll(fileContent); err != nil {
			msg = err.Error()
		}

	} else if typ == "insert" {
		// 若方法为插入但没有传输数据，则错误
		msg = "insert form data error: wrong parameters!"
	}

	// 若删除时既没有 key 也没有 filename
	if len(filename) <= 0 && len(key) <= 0 {
		msg = "invaild parameter for delete operation"
	}

	// 若过程中出现错误
	if len(msg) > 0 {
		context.JSON(400, gin.H{
			"err": msg,
		})
		return
	}

	// 否则按照操作类型进行操作
	switch typ {
	// 插入/更新文件
	case "insert":
		if err := setter.SetData(res, filename, data); err != nil {
			msg = err.Error()
		}
	// 删除条目
	case "delete":
		// 若文件名合法则
		if len(filename) > 0 && filename != "." {
			go setup.DeleteFileData(res, filename)
			database.DeleteFile(database.DataClient, res, filename)
		}
		// 若 key 合法则
		if len(key) > 0 {
			setup.DeleteItem(res, res+"@"+key, 0)
		}
	}

	// 若过程中出现错误
	if len(msg) > 0 {
		context.JSON(400, gin.H{
			"err": msg,
		})
		return
	}

	// 否则返回成功
	context.JSON(200, gin.H{
		"message": "successful!", //result,
	})

}

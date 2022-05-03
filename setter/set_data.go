// @title	SetData
// @description	将前端传来的写入行为描述存储在文件中
// @auth	ryl			2022/4/17		17:00
// @param	resource	string			特型卡片类型（如 "诗词" 和 "车" 等）
// @param	filename	string			文件名
// @param	content		[]ItemSetting	需要写入配置文件的数据
// @return	err			error			错误值

package setter

import (
	"dianasdog/database"
	"dianasdog/getter"
	"dianasdog/setup"
)

func SetData(resource string, filename string, content []byte) error {

	// 插入文件
	if err := database.InsertFile(database.DataClient, resource, filename, content); err != nil {
		return err
	}

	// 查找对应特型卡的配置
	itemSettings, err := getter.GetConfig(resource)
	if err != nil {
		return err
	}

	// 文件拆包（多线程）
	go setup.UnpackXmlData(content, resource, "insert", itemSettings, filename)

	// 无论正确与否都返回 err 的内容
	return nil
}

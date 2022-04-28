// @title	SetConfig
// @description	此函数的用途为，根据数据"类型"写入配置文件，然后根据新旧配置文件的对比结果全量建库。
// @auth	ryl			2022/4/26		16:00
// @param	resource	string			特型卡片类型（如 "诗词" 和 "车" 等）
// @param	content		[]ItemSetting	需要写入配置文件的数据
// @return	err			error			错误值

package setter

import (
	"dianasdog/database"
	"dianasdog/getter"
	"dianasdog/setup"
)

func SetConfig(resource string, content []byte) error {

	// 在数据库中新增此类型
	database.InsertCategory(database.CategoryClient, "word", resource)
	database.CreateCategoryTable(database.CategoryClient, resource)
	database.CreateFileTable(database.DataClient, resource)

	// 新建表格
	database.CreateTableInDict(resource)
	database.CreateTableInPattern(resource)

	// 写入新配置
	err := database.InsertFile(database.ConfigClient, "file", resource, content)
	if err != nil {
		return err
	}

	// 查找新特型卡配置
	newSetting, err := getter.GetConfig(resource)

	// 先删除数据
	go setup.DeleteResData(resource)

	// 再更新数据
	go setup.UpdateResData(resource, newSetting)

	// 无论正确与否都返回 err 的内容
	return err
}

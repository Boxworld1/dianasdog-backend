// @title	SetConfig
// @description	此函数的用途为，根据数据 "类型"，在配置文件中找出对应的 "写入行为"，并反馈相关数据到数据处理函数中。
// @auth	ryl			2022/4/13		16:00
// @param	resource	string			特型卡片类型（如 "诗词" 和 "车" 等）
// @param	content		[]ItemSetting	需要写入配置文件的数据
// @return	err			error			错误值

package setter

import (
	"dianasdog/database"
)

func SetConfig(resource string, content []byte) error {

	// 查找原来特型卡的配置
	// oldSetting, err := getter.GetConfig(resource)
	// if err != nil {
	// 	return err
	// }

	// 写入新配置
	err := database.InsertFile(database.ConfigClient, "file", resource, content)

	// 查找新特型卡配置
	// newSetting, err := getter.GetConfig(resource)

	// 更改对应文件入库的数据（全量建库）

	// 无论正确与否都返回 err 的内容
	return err
}

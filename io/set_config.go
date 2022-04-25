// @title	SetConfig
// @description	此函数的用途为，根据数据 "类型"，在配置文件中找出对应的 "写入行为"，并反馈相关数据到数据处理函数中。
// @auth	ryl			2022/4/13		16:00
// @param	resource	string			特型卡片类型（如 "诗词" 和 "车" 等）
// @param	content		[]ItemSetting	需要写入配置文件的数据
// @return	err			error			错误值

package io

import "dianasdog/database"

func SetConfig(resource string, content []byte) error {

	// 写入配置
	err := database.InsertFile(database.ConfigClient, "file", resource, content)

	// 无论正确与否都返回 err 的内容
	return err
}

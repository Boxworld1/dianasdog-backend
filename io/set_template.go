// @title	SetTemplate
// @description	将前端传来的写入模板配置存储在文件中
// @auth	ryl			2022/4/17		18:00
// @param	resource	string			特型卡片类型（如 "诗词" 和 "车" 等）
// @param	content		[]ItemSetting	需要写入配置文件的数据
// @return	err			error			错误值

package io

import "dianasdog/database"

func SetTemplate(resource string, content []byte) error {

	// 写入配置
	err := database.InsertFile(database.TemplateClient, "file", resource, content)

	// 无论正确与否都返回 err 的内容
	return err
}

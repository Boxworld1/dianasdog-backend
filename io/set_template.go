// @title	SetTemplate
// @description	将前端传来的写入模板配置存储在文件中
// @auth	ryl			2022/4/17		18:00
// @param	resource	string			特型卡片类型（如 "诗词" 和 "车" 等）
// @param	content		[]ItemSetting	需要写入配置文件的数据
// @return	err			error			错误值

package io

import (
	"dianasdog/setup"
	"io/ioutil"
)

func SetTemplate(resource string, content []byte) error {

	// 得到此文件的绝对路径
	abspath, _ := setup.GetAbsPath()

	// 查找对应类型的 template 文档路径
	filepath := abspath + "template/" + resource + ".json"

	// 写入配置
	err := ioutil.WriteFile(filepath, content, 0644)

	// 无论正确与否都返回 err 的内容
	return err
}

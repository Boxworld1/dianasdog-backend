// @title	SetData
// @description	将前端传来的写入行为描述存储在文件中
// @auth	ryl			2022/4/17		17:00
// @param	resource	string			特型卡片类型（如 "诗词" 和 "车" 等）
// @param	filename	string			文件名
// @param	content		[]ItemSetting	需要写入配置文件的数据
// @return	err			error			错误值

package io

import (
	"dianasdog/path"
	"io/ioutil"
	"os"
)

func SetData(resource string, filename string, content []byte) error {

	// 得到此文件的绝对路径
	abspath, _ := path.GetAbsPath()

	// 查找对应类型的文件路径（先记为 .txt）
	filepath := abspath + "data/" + resource + "/"
	tmppath := filepath + "1.txt"

	// 新建文件夹
	_ = os.MkdirAll(filepath, os.ModePerm)

	// 写入配置
	err := ioutil.WriteFile(tmppath, content, 0644)

	// 写入后改回 .xml
	os.Rename(tmppath, filepath+filename)

	// 无论正确与否都返回 err 的内容
	return err
}

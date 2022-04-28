// @title	SetTemplate
// @description	将前端传来的写入模板配置存储在文件中
// @auth	ryl			2022/4/29	1:00
// @param	resource	string		特型卡片类型（如 "诗词" 和 "车" 等）
// @param	content		[]string	需要写入配置文件的数据
// @param	opType		string		插入或删除
// @return	err			error		错误值

package setter

import "dianasdog/database"

func SetTemplate(resource string, content []string, opType string) error {

	var err error
	switch opType {
	case "insert":
		// 写入配置
		for _, data := range content {
			err = database.InsertToPattern(resource, data)
		}
	case "delete":
		// 删除配置
		for _, data := range content {
			err = database.DeleteFromPattern(resource, data)
		}
	}

	// 无论正确与否都返回 err 的内容
	return err
}

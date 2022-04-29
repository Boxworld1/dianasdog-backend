// @title	SetTemplate
// @description	将前端传来的写入模板配置存储在文件中
// @auth	ryl			2022/4/29	1:00
// @param	resource	string		特型卡片类型（如 "诗词" 和 "车" 等）
// @param	content		[]string	需要写入配置文件的数据
// @param	opType		string		插入或删除
// @return	err			error		错误值

package setter

import "dianasdog/database"

func SetTemplateAll(content []string, opType string) error {
	var err error
	// 取出所有类型
	data, _ := database.GetAllCategory(database.CategoryClient, "word")
	// 按类別插入
	for _, res := range data {
		err = SetTemplate(res, content, opType)
	}
	return err
}

func SetTemplate(resource string, content []string, opType string) error {
	var err error

	// 若需全部插入则
	if resource == "all" {
		return SetTemplateAll(content, opType)
	}

	// 否则若为单一类型
	switch opType {
	case "insert":
		// 写入配置
		for _, data := range content {
			if err = database.InsertToPattern(resource, data); err != nil {
				return err
			}
		}
	case "delete":
		// 删除配置
		for _, data := range content {
			if err = database.DeleteFromPattern(resource, data); err != nil {
				return err
			}
		}
	}

	// 无论正确与否都返回 err 的内容
	return nil
}

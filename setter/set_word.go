// @title	SetWord
// @description	将前端传来的词存储在文件中
// @auth	ryl			2022/4/29		1:00
// @param	resource	string			特型卡片类型（如 "诗词" 和 "车" 等）
// @param	content		[]string	需要写入配置文件的数据
// @return	err			error			错误值

package setter

import (
	"dianasdog/database"
	"fmt"
)

func GetDocid(resource string, wordType string, data string) string {
	return resource + "@" + wordType + "@" + data
}

func SetWordAll(content []string, opType string, wordType string) error {
	var err error
	// 取出所有类型
	data, _ := database.GetAllCategory(database.CategoryClient, "word")
	// 按类別插入
	for _, res := range data {
		err = SetWord(res, content, opType, wordType)
	}
	return err
}

func SetWord(resource string, content []string, opType string, wordType string) error {

	var err error
	// 若需全部插入则
	if resource == "all" {
		return SetWordAll(content, opType, wordType)
	}

	// 否则若为单一类型
	switch opType {
	case "insert":
		// 若为插入则插入词
		for _, data := range content {
			fmt.Println("insert " + data + " as " + wordType)
			if err = database.InsertToDict(resource, GetDocid(resource, wordType, data), wordType, data); err != nil {
				return err
			}
		}
	case "delete":
		// 删除词
		for _, data := range content {
			fmt.Println("delete " + data + " as " + wordType)
			if err = database.DeleteByDocidFromDict(resource, GetDocid(resource, wordType, data)); err != nil {
				return err
			}
		}
	}

	// 返回无错误
	return nil
}

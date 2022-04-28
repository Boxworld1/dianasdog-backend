// @title	SetWord
// @description	将前端传来的词存储在文件中
// @auth	ryl			2022/4/29		1:00
// @param	resource	string			特型卡片类型（如 "诗词" 和 "车" 等）
// @param	content		[]string	需要写入配置文件的数据
// @return	err			error			错误值

package setter

import "dianasdog/database"

func GetDocid(resource string, wordType string, data string) string {
	return resource + "@" + wordType + "@" + data
}

func SetWord(resource string, content []string, opType string, wordType string) error {

	var err error
	switch opType {
	case "insert":
		// 若为插入则插入词
		for _, data := range content {
			err = database.InsertToDict(resource, GetDocid(resource, wordType, data), wordType, data)
		}
	case "delete":
		// 删除词
		for _, data := range content {
			err = database.DeleteByDocidFromDict(resource, GetDocid(resource, wordType, data))
		}
	}

	// 无论正确与否都返回 err 的内容
	return err
}

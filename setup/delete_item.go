// @title	DeleteItem
// @description	此函数的用途为删除数据
// @auth	ryl		2022/4/28		14:20
// @param	resource	string	特型卡id
// @param	docid		string	docid
// @param	opType		int		操作类型(全局删除:0, 更新:1)
// @return	err			error	错误值

package setup

import "dianasdog/database"

func DeleteItem(resource string, docid string, opType int) error {

	// 在字典中删除数据
	database.DeleteByDocidFromDict(resource, docid)
	// 在 redis 中删除数据
	database.DeleteFromRedis(database.RedisClient, docid)
	// 在 ES 中删除数据
	database.DeleteFromES(resource, database.EsClient, docid)
	// 在数据记录中删除数据
	if opType == 0 {
		database.DeleteFile(database.DocidClient, resource, docid)
	}

	return nil
}

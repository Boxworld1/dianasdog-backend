// @title	TestSearch
// @description	此函数的用途为检查 Search 函数的正确性
// @auth	ryl		2022/4/25	22:15
// @param	t		*testing.T	testing 用参数

package search

import (
	"dianasdog/database"
	"dianasdog/setter"
	"testing"
)

func TestSearch(t *testing.T) {

	// 插入必要的测试数据
	setter.SetWord("all", []string{"的", "是多少", "是什么", "有什么"}, "insert", "garbage")
	setter.SetWord("test", []string{"价格", "售价"}, "insert", "intent")
	setter.SetTemplate("test", []string{"title+garbage+intent+garbage"}, "insert")
	database.InsertToEs("test", database.EsClient, "test@宝马", "宝马")
	database.SetToRedis(database.RedisClient, "test@宝马", "{\"Test\": \"测试成功\"}")

	// 搜索测试
	result := Search("白居易")

	if result == nil {
		t.Error("搜索失败")
	}
}

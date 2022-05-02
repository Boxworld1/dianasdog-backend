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
	database.SetToRedis(database.RedisClient, "test@宝马", `{"type":"car","item":{"display":{"title":"\u534E\u6668\u5B9D\u9A6C\u65B0\u80FD\u6E90","sub_brands":{"item":[{"tag":"\u70ED\u95E8","series_list":{"item":[{"hot":"46844","price":"39.99\u4E07\u8D77","series_name":"\u5B9D\u9A6CiX3","series_brand":"\u534E\u6668\u5B9D\u9A6C"},{"hot":"36609","price":"49.99\u4E07\u8D77","series_name":"\u5B9D\u9A6C5\u7CFBPHEV","series_brand":"\u534E\u6668\u5B9D\u9A6C"},{"hot":"8938","price":"39.98\u4E07\u8D77","series_name":"\u5B9D\u9A6CX1 PHEV","series_brand":"\u534E\u6668\u5B9D\u9A6C"},{"hot":"6957","price":"\u672A\u4E0A\u5E02","series_name":"\u5B9D\u9A6Ci3","series_brand":"\u534E\u6668\u5B9D\u9A6C"}]}},{"series_list":{"item":[{"hot":"36609","price":"49.99\u4E07\u8D77","series_name":"\u5B9D\u9A6C5\u7CFBPHEV","series_brand":"\u534E\u6668\u5B9D\u9A6C"},{"hot":"6957","price":"\u672A\u4E0A\u5E02","series_name":"\u5B9D\u9A6Ci3","series_brand":"\u534E\u6668\u5B9D\u9A6C"}]},"tag":"\u8F7F\u8F66"},{"tag":"SUV","series_list":{"item":[{"hot":"46844","price":"39.99\u4E07\u8D77","series_name":"\u5B9D\u9A6CiX3","series_brand":"\u534E\u6668\u5B9D\u9A6C"},{"hot":"8938","price":"39.98\u4E07\u8D77","series_name":"\u5B9D\u9A6CX1 PHEV","series_brand":"\u534E\u6668\u5B9D\u9A6C"}]}}]}}},"picture":["http:\/\/p1-dcd.byteimg.com\/img\/motor-img\/6ed2508e2edf0fc8de32cb00897d2271~240x0.png","http:\/\/p1-dcd.byteimg.com\/img\/motor-img\/59653d417fe0fbb72328343b169e8d98~240x0.png","http:\/\/p6-dcd.byteimg.com\/img\/motor-img\/2a738cde9d6b91bd5fd67606fc3b308b~240x0.png","http:\/\/p1-dcd.byteimg.com\/img\/motor-img\/73765fe41507955301bd91c390a0505b~240x0.png","http:\/\/p3-dcd.byteimg.com\/img\/motor-img\/59653d417fe0fbb72328343b169e8d98~240x0.png","http:\/\/p3-dcd.byteimg.com\/img\/motor-img\/73765fe41507955301bd91c390a0505b~240x0.png","http:\/\/p1-dcd.byteimg.com\/img\/motor-img\/6ed2508e2edf0fc8de32cb00897d2271~240x0.png","http:\/\/p9-dcd.byteimg.com\/img\/motor-img\/2a738cde9d6b91bd5fd67606fc3b308b~240x0.png"]}`)

	// 搜索测试
	result := Search("白居易")

	if result == nil {
		t.Error("搜索失败")
	}
}

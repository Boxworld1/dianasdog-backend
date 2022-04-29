package search

import (
	"dianasdog/database"
	"testing"
)

func TestQueryUnderstanding(t *testing.T) {
	database.CreateTableInPattern("test")
	err := database.InsertToPattern("test", "title+garbage+intent+garbage")
	if err != nil {
		t.Error("插入失败", err)
	}

	// 测试鲁棒性用
	err = database.InsertToPattern("test", "tilte+garbage+intent+garbage")
	if err != nil {
		t.Error("插入失败", err)
	}

	database.CreateTableInDict("test")
	database.InsertToDict("test", "10086", "title", "宝马")
	database.InsertToDict("test", "0", "garbage", "的")
	database.InsertToDict("test", "0", "garbage", "是多少")
	database.InsertToDict("test", "1", "intent", "价格")

	QueryUnderstanding([]string{"test"}, "宝马的价格是多少")
	// if res[0].detail[0] != "宝马" || res[0].detail[1] != "的" || res[0].detail[2] != "价格" || res[0].detail[3] != "是多少" {
	// 	t.Error("wrong answer")
	// }
}

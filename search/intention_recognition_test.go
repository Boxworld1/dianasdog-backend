// @title	TestIntentionRecognition
// @description	检查意图识別功能
// @auth	jz	2022/4/7	12:00
// @param	t	*testing.T	testing 用参数

package search

import (
	"dianasdog/database"
	"testing"
)

func TestIntentionRecognition(t *testing.T) {
	database.CreateTableInDict("test")
	database.InsertToDict("test", "10086", "title", "宝马")
	dict := IntentionRecognition("宝马的价格是多少")
	if len(dict) == 0 {
		t.Error("invalid length")
	}
}

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
	} else {
		if dict[0] != "test" {
			t.Error("wrong answer")
		}
	}
}

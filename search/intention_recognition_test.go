package search

import (
	"dianasdog/database"
	"testing"
)

func TestIntentionRecognition(t *testing.T) {
	d := []string{"id", "title"}
	database.CreateTableFromDict(database.DictClient, "test", d)
	item := []string{"test0", "宝马"}
	database.InsertToDict(database.DictClient, "test", item)
	dict := IntentionRecognition("宝马的价格是多少")
	if len(dict) == 0 {
		t.Error("invalid length")
	} else {
		if dict[0] != "test" {
			t.Error("wrong answer")
		}
	}
}

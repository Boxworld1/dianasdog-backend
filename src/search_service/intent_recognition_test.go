// @title		TestIntentRecognition
// @description	此函数的用途为检查 IntentRecognition 函数的正确性
// @auth		ryl				2022/3/24		11:00
// @param		t				*testing.T		testing 用参数

package search_service

import (
	"testing"
)

func TestIntentRecognition(t *testing.T) {

	resources := IntentRecognition("carcdc")

	// 配置数量错误
	if len(resources) != 2 {
		t.Error("配置数量有误")
	} else {
		item0 := resources[0]
		item1 := resources[1]
		// 配置错误
		if item0 != "car" || item1 != "cd" {
			t.Error("回传类型错误")
		}
	}

	resources = IntentRecognition("I have a dream!")
	if len(resources) != 0 {
		t.Error("配置数量有误")
	}

	resources = IntentRecognition("I want to buy a cargo.")
	if len(resources) != 1 {
		t.Error("配置数量有误")
	} else {
		item0 := resources[0]
		// 配置错误
		if item0 != "car" {
			t.Error("回传类型错误")
		}
	}
}

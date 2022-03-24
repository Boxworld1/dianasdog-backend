// @title		TestWordSearch
// @description	此函数的用途为检查 GetConfig 函数的正确性
// @auth		ryl				2022/3/24		10:00
// @param		t				*testing.T		testing 用参数

package search_service

import (
	"testing"
)

func TestWordSearch(t *testing.T) {

	var resources = make([]string, 0)
	result := WordSearch("car")
	resources = append(resources, result...)

	// 配置数量错误
	if len(resources) != 1 {
		t.Error("配置数量有误")
	} else {
		item0 := resources[0]
		// 配置错误
		if item0 != "car" {
			t.Error("回传类型错误")
		}
	}

	result = WordSearch("cd")
	resources = append(resources, result...)

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

	result = WordSearch("hello")
	resources = append(resources, result...)

	// 配置数量错误
	if len(resources) != 2 {
		t.Error("配置数量有误")
	}

	result = WordSearch("cad")
	resources = append(resources, result...)
	if len(resources) != 3 {
		t.Error("配置数量有误")
	}
}

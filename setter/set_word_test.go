// @title	TestSetWord
// @description	此函数的用途为检查 SetWord 函数的正确性
// @auth	ryl		2022/4/29	1:00
// @param	t		*testing.T	testing 用参数

package setter

import (
	"dianasdog/testcase"
	"testing"
)

func TestSetWord(t *testing.T) {

	// 初始化测例
	if err := testcase.SetTestData(0); err != nil {
		t.Error("测例建造失败")
	}

	// 插入存在的特型卡中
	err := SetWord("testdata", []string{"intent"}, "insert", "intent")
	// 测试时出错
	if err != nil {
		t.Error(err)
	}

	// 插入所有存在的特型卡中
	err = SetWord("all", []string{"intent"}, "insert", "intent")
	// 测试时出错
	if err != nil {
		t.Error(err)
	}

	// 从特型卡中删除
	err = SetWord("testdata", []string{"intent"}, "delete", "intent")
	// 测试时出错
	if err != nil {
		t.Error(err)
	}

	// 从所有特型卡中删除
	err = SetWord("all", []string{"intent"}, "delete", "intent")
	// 测试时出错
	if err != nil {
		t.Error(err)
	}
}

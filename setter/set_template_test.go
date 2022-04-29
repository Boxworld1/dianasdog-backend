// @title	TestSetTemplate
// @description	此函数的用途为检查 SetTemplate 函数的正确性
// @auth	ryl		2022/4/13	10:00
// @param	t		*testing.T	testing 用参数

package setter

import (
	"dianasdog/database"
	"dianasdog/testcase"
	"testing"
)

func TestSetTemplate(t *testing.T) {

	// 初始化测例
	if err := testcase.SetTestData(0); err != nil {
		t.Error("测例建造失败")
	}

	// 读入文件
	data, err := database.FetchAllPattern("testdata")
	if err != nil {
		t.Error("测试文件有误")
	}

	// 插入存在的特型卡中
	err = SetTemplate("testdata", data, "insert")
	// 测试时出错
	if err != nil {
		t.Error(err)
	}

	// 插入所有存在的特型卡中
	err = SetTemplate("all", data, "insert")
	// 测试时出错
	if err != nil {
		t.Error(err)
	}

	// 从特型卡中删除
	err = SetTemplate("testdata", data, "delete")
	// 测试时出错
	if err != nil {
		t.Error(err)
	}

	// 从所有存在的特型卡中删除
	err = SetTemplate("all", data, "delete")
	// 测试时出错
	if err != nil {
		t.Error(err)
	}
}

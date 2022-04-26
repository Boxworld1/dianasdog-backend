// @title	TestSetConfig
// @description	此函数的用途为检查 SetConfig 函数的正确性
// @auth	ryl		2022/4/13	10:00
// @param	t		*testing.T	testing 用参数
package setter

import (
	"dianasdog/database"
	"dianasdog/testcase"
	"testing"
)

func TestSetConfig(t *testing.T) {

	// 初始化测例
	if err := testcase.SetTestData(0); err != nil {
		t.Error("测例建造失败")
	}

	// 读入文件
	data, err := database.GetFile(database.ConfigClient, "file", "testdata")
	if err != nil {
		t.Error("测试文件有误")
	}

	// 初始化测例
	if err := testcase.SetTestData(1); err != nil {
		t.Error("测例建造失败")
	}

	err = SetConfig("testdata", data)
	// 测试时出错
	if err != nil {
		t.Error(err)
	}

}

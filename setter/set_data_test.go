// @title	TestSetData
// @description	此函数的用途为检查 SetData 函数的正确性
// @auth	ryl		2022/4/14	10:30
// @param	t		*testing.T	testing 用参数

package setter

import (
	"dianasdog/database"
	"dianasdog/testcase"
	"testing"
)

func TestSetData(t *testing.T) {

	// 初始化测例
	if err := testcase.SetTestData(); err != nil {
		t.Error("测例建造失败")
	}

	// 读入文件
	data, err := database.GetFile(database.DataClient, "testdata", "testcase.xml")
	if err != nil {
		t.Error("测试文件有误")
	}

	err = SetData("testdata", "testcase.xml", data)
	// 测试时出错
	if err != nil {
		t.Error(err)
	}

}

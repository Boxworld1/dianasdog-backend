// @title	TestUpdateResData
// @description	此函数的用途为检查 UpdateResData 函数的正确性
// @auth	ryl		2022/4/26	15:10
// @param	t		*testing.T	testing 用参数

package setup

import (
	"dianasdog/database"
	"dianasdog/getter"
	"dianasdog/testcase"
	"testing"
)

func TestUpdateResData(t *testing.T) {

	// 初始化测例
	if err := testcase.SetTestData(0); err != nil {
		t.Error("测例建造失败")
	}

	// 查找特型卡配置
	itemSettings, err := getter.GetConfig("testdata")
	if err != nil {
		t.Error("读入配置出错")
	}

	// 测试 SaveItem
	database.DropCategory(database.CategoryClient, "testdata")
	database.CreateCategoryTable(database.CategoryClient, "testdata")

	// 全量建库测试
	err = UpdateResData("testdata", itemSettings)
	if err != nil {
		t.Error("检测到不存在的错误！")
	}

	// 全量建库：不存在特型卡
	err = UpdateResData("testdata_apple", itemSettings)
	if err == nil {
		t.Error("检测不到存在的错误！")
	}

}

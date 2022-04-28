// @title	TestDeleteResData
// @description	此函数的用途为检查 DeleteResData 函数的正确性
// @auth	ryl		2022/4/28	15:10
// @param	t		*testing.T	testing 用参数

package setup

import (
	"dianasdog/testcase"
	"testing"
)

func TestDeleteResData(t *testing.T) {

	// 初始化测例
	if err := testcase.SetTestData(0); err != nil {
		t.Error("测例建造失败")
	}

	// 全量删库测试
	err := DeleteResData("testdata")
	if err != nil {
		t.Error("检测到不存在的错误！")
	}

	// 全量删库：不存在特型卡
	err = DeleteResData("testdata_apple")
	if err == nil {
		t.Error("检测不到存在的错误！")
	}

}

// @title		TestAddData
// @description	测试 AddData 函数的正确性
// @auth		ryl				2022/4/7		10:00
// @param		t				*testing.T		testing 用参数

package setup

import (
	"testing"
)

func TestAddData(t *testing.T) {
	// 读入错误/无文件路径
	err := AddData("test", "test")
	if err == nil {
		t.Error("读入错误文件")
	}

	// 读入测试文件（是否有误由下层函数判断）
	err = AddData("../data/testcase/", "test")
	if err != nil {
		t.Error(err)
	}
}

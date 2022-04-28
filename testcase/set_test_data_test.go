// @title	TestSetTestData
// @description	此函数的用途为检查 SetTestData 函数的正确性
// @auth	ryl		2022/4/26	10:30
// @param	t		*testing.T	testing 用参数

package testcase

import (
	"testing"
)

func TestSetTestData(t *testing.T) {

	for i := 0; i < 2; i++ {
		// 初始化测例
		err := SetTestData(i)

		// 测试时出错
		if err != nil {
			t.Error(err)
		}
	}

}

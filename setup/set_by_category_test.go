// @title		TestSetByCategory
// @description	测试 TestSetByCategory 的功能
// @auth		ryl				2022/4/7		10:00
// @return		err				error			错误值

package setup

import (
	"testing"
)

func TestSetByCategory(t *testing.T) {
	// 由于下层函数已经判断了读入内容的正确性，因此这里仅作接口读入操作
	err := SetByCategory()
	if err != nil {
		t.Error(err)
	}
}

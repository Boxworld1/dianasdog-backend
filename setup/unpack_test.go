// @title		SetupTest
// @description	此函数的用途为检查 setup 包的正确性
// @auth		ryl				2022/3/17		11:05
// @param		t				*testing.T		testing 用参数

package setup

import (
	"testing"
)

func TestUnpackXml(t *testing.T) {

	// 查找不存在的文件
	err := UnpackXmlFile("apple.xml", "testcase_apple")
	if err == nil {
		t.Error("无法检测问题，错误！")
	}

	// 查找存在的特型卡配置
	// err = UnpackXmlFile("testcase.xml", "testcase")
	// if err != nil {
	// 	t.Error("检测到不存在的错误！")
	// }

}

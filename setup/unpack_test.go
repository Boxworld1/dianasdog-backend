// @title	TestUnpackXml
// @description	此函数的用途为检查 UnpackXml 函数的正确性
// @auth	ryl		2022/4/25	17:05
// @param	t		*testing.T	testing 用参数

package setup

import (
	"dianasdog/io"
	"testing"
)

func TestUnpackXml(t *testing.T) {

	// 初始化测例
	if err := io.SetTestData(); err != nil {
		t.Error("测例建造失败")
	}

	// 查找不存在的文件
	err := UnpackXmlFile("apple.xml", "testcase_apple")
	if err == nil {
		t.Error("无法检测问题，错误！")
	}

	// 查找存在的特型卡配置
	err = UnpackXmlFile("testcase.xml", "testdata")
	if err != nil {
		t.Error("检测到不存在的错误！")
	}

}

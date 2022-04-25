// @title		TestStoreItem
// @description	此函数的用途为检查 setup 包 StoreItem 的正确性
// @auth		ryl				2022/3/17		11:05
// @param		t				*testing.T		testing 用参数

package setup

import (
	"dianasdog/io"
	"dianasdog/path"
	"testing"

	"github.com/beevik/etree"
)

func TestStoreItem(t *testing.T) {
	// 得到此文件的绝对路径
	abspath, _ := path.GetAbsPath()
	path := abspath + "data/testcase/testcase_normal.xml"
	doc := etree.NewDocument()

	// 读入文件错误
	if err := doc.ReadFromFile(path); err != nil {
		t.Error(err)
	}

	root := doc.SelectElement("DOCUMENT")

	// 查找存在的特型卡配置
	itemSetting, err := io.GetConfig("testcase_poem")
	if err == nil {
		t.Error("无法检测问题，错误！")
	}

	// 插入正常数据
	for _, item := range root.SelectElements("item") {
		err := StoreItem(item, "testcase_poem", "insert", "0", itemSetting)
		if err != nil {
			t.Error(err)
		}
	}
}

// @title		SetupTest
// @description	此函数的用途为检查 setup 包的正确性
// @auth		ryl				2022/3/17		11:05
// @param		t				*testing.T		testing 用参数

package setup

import (
	"dianasdog/path"
	"testing"

	"github.com/beevik/etree"
)

func SetupTest(t *testing.T) {
	// 得到此文件的绝对路径
	abspath, _ := path.GetAbsPath()

	path := abspath + "data/testcase/testcase_normal.xml"
	doc := etree.NewDocument()

	// 读入文件错误
	if err := doc.ReadFromFile(path); err != nil {
		t.Error(err)
	}

	root := doc.SelectElement("DOCUMENT")

	// 插入不存在的特型卡类型
	testItem := root.FindElement("item")
	if testItem != nil {
		myErr := StoreItem(testItem, "apple", "delete", "0")
		if myErr == nil {
			t.Error("无法检测问题，错误！")
		}
	}

	// 插入正常数据
	// for _, item := range root.SelectElements("item") {
	// 	err := StoreItem(item, "testcase_poem", "insert", "0")
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// }
}

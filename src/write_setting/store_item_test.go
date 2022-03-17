// @title		TestStoreItem
// @description	此函数的用途为检查 StoreItem 函数的正确性
// @auth		ryl				2022/3/17		11:05
// @param		t				*testing.T		testing 用参数

package write_setting

import (
	"testing"

	"github.com/beevik/etree"
)

func TestStoreItem(t *testing.T) {
	path := "./testcase/test.xml"
	doc := etree.NewDocument()

	// 读入文件错误
	if err := doc.ReadFromFile(path); err != nil {
		t.Error(err)
	}

	root := doc.SelectElement("DOCUMENT")

	// 插入不存在的特型卡类型
	testItem := root.FindElement("item")
	if testItem != nil {
		myErr := StoreItem(testItem, "apple", "delete")
		if myErr == nil {
			t.Error("无法检测问题，错误！")
		}
	}

	// 插入正常数据
	for _, item := range root.SelectElements("item") {
		err := StoreItem(item, "poem", "insert")
		if err != nil {
			t.Error(err)
		}
	}
}

// @title		TestStoreItem
// @description	此函数的用途为检查 StoreItem 函数的正确性
// @auth		ryl				2022/4/25		18:05
// @param		t				*testing.T		testing 用参数

package setup

import (
	"dianasdog/database"
	"dianasdog/getter"
	"dianasdog/testcase"
	"testing"

	"github.com/beevik/etree"
)

func TestStoreItem(t *testing.T) {

	// 初始化测例
	if err := testcase.SetTestData(); err != nil {
		t.Error("测例建造失败")
	}

	// 读入文件
	data, err := database.GetFile(database.DataClient, "testdata", "testcase.xml")
	if err != nil {
		t.Error("测试文件有误")
	}

	// 读入文件错误
	doc := etree.NewDocument()
	if err := doc.ReadFromString(string(data)); err != nil {
		t.Error(err)
	}

	root := doc.SelectElement("DOCUMENT")

	// 查找特型卡配置
	itemSetting, err := getter.GetConfig("testdata")
	if err != nil {
		t.Error("无法检测问题，错误！")
	}

	// 插入正常数据
	for _, item := range root.SelectElements("item") {
		err := StoreItem(item, "testdata", "insert", "0", itemSetting)
		if err != nil {
			t.Error(err)
		}
	}
}

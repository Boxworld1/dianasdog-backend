// @title		TestStoreItem
// @description	此函数的用途为检查 setup 包 StoreItem 的正确性
// @auth		ryl				2022/3/17		11:05
// @param		t				*testing.T		testing 用参数

package setup

import (
	"dianasdog/database"
	"dianasdog/io"
	"testing"

	"github.com/beevik/etree"
)

func TestStoreItem(t *testing.T) {

	// 初始化测例
	if err := io.SetTestData(); err != nil {
		t.Error("测例建造失败")
	}

	// 读入文件
	data, err := database.GetFile(database.DataClient, "testcase", "testcase.xml")
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
	itemSetting, err := io.GetConfig("testcase")
	if err != nil {
		t.Error("无法检测问题，错误！")
	}

	// 插入正常数据
	for _, item := range root.SelectElements("item") {
		err := StoreItem(item, "testcase", "insert", "0", itemSetting)
		if err != nil {
			t.Error(err)
		}
	}
}

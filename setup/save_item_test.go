// @title	TestSaveItem
// @description	此函数的用途为检查 SaveItem 函数的正确性
// @auth	ryl		2022/4/28	12:05
// @param	t		*testing.T	testing 用参数

package setup

import (
	"dianasdog/database"
	"dianasdog/testcase"
	"testing"

	"github.com/beevik/etree"
)

func TestSaveItem(t *testing.T) {

	// 初始化测例
	if err := testcase.SetTestData(0); err != nil {
		t.Error("测例建造失败")
	}

	// 读入文件
	data, err := database.GetFile(database.DataClient, "testdata", "testcase.xml")
	if err != nil {
		t.Error("测试文件有误")
	}

	// 读入文件错误
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		t.Error(err)
	}

	root := doc.SelectElement("DOCUMENT")

	// 插入正常数据
	for _, item := range root.SelectElements("item") {
		err := SaveItem(item, "testdata")
		if err != nil {
			t.Error(err)
		}
		break
	}

	cnt, _ := database.CountCategory(database.CategoryClient, "testdata")

	// 若数据个数不足，则错误
	if cnt < 9 {
		t.Error("测例有误")
	}
}

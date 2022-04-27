// @title   UnpackXmlData
// @description 拆解 XML 并入库
// @auth	ryl			2022/4/26	13:30
// @param	data		[]byte		数据
// @param	resource	string		特型卡类型
// @param	opType		string		操作类型（insert/delete）
// @param	itemSettings	[]getter.ItemSetting	写入行为
// @return  err			error		non-nil when fileName is wrong

package setup

import (
	"dianasdog/database"
	"dianasdog/getter"

	"github.com/beevik/etree"
)

func UnpackXmlData(data []byte, resource string, opType string, itemSettings []getter.ItemSetting) error {

	// 将数据放入 etree 中
	doc := etree.NewDocument()
	err := doc.ReadFromBytes(data)

	// 文件有误
	if err != nil {
		return err
	}

	// 新建表格
	database.CreateFileTable(database.DocidClient, resource)

	// 按 item 划分 etree
	root := doc.SelectElement("DOCUMENT")
	itemList := root.SelectElements("item")

	// 遍历所有 item 并存入数据库
	for _, item := range itemList {
		docid := GetDocid(item, resource)

		// 将数据存入 docid 库
		itemDoc := etree.NewDocument()
		itemDoc.SetRoot(item)
		str, _ := itemDoc.WriteToBytes()
		database.InsertFile(database.DocidClient, resource, docid, str)

		// 然后存入其他数据库
		if err := StoreItem(item, resource, opType, docid, itemSettings); err != nil {
			return err
		}
	}

	return nil
}

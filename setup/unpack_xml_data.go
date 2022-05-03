// @title   UnpackXmlData
// @description 拆解 XML 并入库
// @auth	ryl			2022/4/26	13:30
// @param	data		[]byte		数据
// @param	resource	string		特型卡类型
// @param	opType		string		操作类型（insert/delete）
// @param	itemSettings	[]getter.ItemSetting	写入行为
// @param filename	string		文件名
// @return  err			error		non-nil when fileName is wrong

package setup

import (
	"dianasdog/database"
	"dianasdog/getter"

	"github.com/beevik/etree"
)

func UnpackXmlData(data []byte, resource string, opType string, itemSettings []getter.ItemSetting, filename string) error {

	// 将数据放入 etree 中
	doc := etree.NewDocument()
	err := doc.ReadFromBytes(data)

	// 文件有误
	if err != nil {
		return err
	}

	// 新建表格
	database.CreateDocidTable(database.DocidClient, resource)

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
		database.InsertDocid(database.DocidClient, resource, docid, str, filename)

		// 检查是否有此特型卡
		cnt, err := database.CountCategory(database.CategoryClient, resource)
		if err != nil {
			return err
		}

		// 若相关类型未存入过 item
		if cnt == 0 {
			SaveItem(item, resource)
		}

		// 然后存入其他数据库
		StoreItem(item, resource, docid, itemSettings)

	}

	return nil
}

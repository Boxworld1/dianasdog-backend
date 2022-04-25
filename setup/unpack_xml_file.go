// @title   UnpackXmlFile
// @description  unpack large XML file to single item
// @auth    于沛楠       2022/3/16
// @auth	ryl			2022/4/20	22:00
// @param	filename     string             the name of XML File to unpack
//			resource	 string				the category name of special card
// @return  itemList     []*etree.Element   XML <item> array (the itemList uses etree from "github.com/beevik/etree")
//			itemCount    int                total count of <item></item>
//          docIDList    []int				docID of each item in itemList
//			resourceName string				the category name of special card
//		    err          error              non-nil when fileName is wrong

package setup

import (
	"dianasdog/database"

	"github.com/beevik/etree"
)

func UnpackXmlFile(filename string, resource string) error {

	// 取得 xml 数据
	data, err := database.GetFile(database.DataClient, resource, filename)
	if err != nil {
		return err
	}

	// 将数据放入 etree 中
	doc := etree.NewDocument()
	err = doc.ReadFromString(string(data))

	// 文件名有误
	if err != nil {
		return err
	}

	// 按 item 划分 etree
	root := doc.SelectElement("DOCUMENT")
	itemList := root.SelectElements("item")

	// 查找对应特型卡的配置
	itemSettings, err := GetConfig(resource)
	if err != nil {
		return err
	}

	// 遍历所有 item 并存入数据库
	for _, item := range itemList {
		docid := GetDocid(item, resource)
		if err := StoreItem(item, resource, "insert", docid, itemSettings); err != nil {
			return err
		}
	}

	return nil
}

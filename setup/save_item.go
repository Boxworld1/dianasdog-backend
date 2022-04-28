// @title	SaveItem
// @description	此函数的用途为得到数据的所有子结点路径
// @auth	ryl		2022/4/28		12:05
// @param	item	*etree.Element	数据 item
// @param	resource	string		特型卡 Id
// @return	docid	string			docid

package setup

import (
	"dianasdog/database"
	"strings"

	"github.com/beevik/etree"
)

func SaveItem(parent *etree.Element, resource string) error {
	for _, item := range parent.ChildElements() {
		if item.ChildElements() == nil {
			// 得到绝对路径
			str := item.GetPath()
			// 去掉句首的 "/"
			str = strings.Replace(str, "/", "", 1)
			// 去掉句首的 "DOCUMENT/"
			str = strings.Replace(str, "DOCUMENT/", "", 1)
			// 去掉句首的 "item/"
			str = strings.Replace(str, "item/", "", 1)
			// 将所有斜线转为 "."
			str = strings.Replace(str, "/", ".", -1)
			// 结果存入数据库
			database.InsertCategory(database.CategoryClient, resource, str)
		} else {
			SaveItem(item, resource)
		}
	}
	return nil
}

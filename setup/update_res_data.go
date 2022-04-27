// @title   UpdateResData
// @description 更新特定特型卡数据（全量建库）
// @auth	ryl			2022/4/26	22:00
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

func UpdateResData(resource string, opType string, itemSettings []getter.ItemSetting) error {

	// 查找特型卡类型下的所有数据
	data, err := database.GetAllFile(database.DocidClient, resource)

	// 若特型卡类型错误
	if err != nil {
		return err
	}

	// 数据分解
	for _, block := range data {
		// 取得 docid 和对应内容
		docid := block.Name
		content := block.Data

		// 将其转化为 etree 方便写入
		doc := etree.NewDocument()
		if err := doc.ReadFromBytes(content); err != nil {
			return err
		}
		item := doc.Root()

		// 直接调用 StoreItem 储存数据
		StoreItem(item, resource, opType, docid, itemSettings)
	}

	return nil
}

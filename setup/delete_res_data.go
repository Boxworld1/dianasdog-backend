// @title   DeleteResData
// @description 更新特定特型卡数据（全量建库）
// @auth	ryl			2022/4/26	22:00
// @param	resource	string		特型卡类型
// @return  err			error		non-nil when fileName is wrong

package setup

import (
	"dianasdog/database"
)

func DeleteResData(resource string) error {

	// 查找特型卡类型下的所有数据
	data, err := database.GetAllFile(database.DocidClient, resource)

	// 若特型卡类型错误
	if err != nil {
		return err
	}

	// 数据分解
	for _, block := range data {
		// 取得 docid
		docid := block.Name
		DeleteItem(resource, docid)
	}

	return nil
}

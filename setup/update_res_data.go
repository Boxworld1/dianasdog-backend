// @title   UpdateResData
// @description 更新特定特型卡数据（全量建库）
// @auth	ryl			2022/4/26	22:00
// @param	resource	string		特型卡类型
// @param	type_		string		操作类型（insert/delete）
// @param	itemSettings	[]getter.ItemSetting	写入行为
// @return  err			error		non-nil when fileName is wrong

package setup

import (
	"dianasdog/database"
	"dianasdog/getter"
)

func UpdateResData(resource string, type_ string, itemSettings []getter.ItemSetting) error {

	filenames, _ := database.GetFileName(database.DataClient, resource)

	for _, file := range filenames {
		go UnpackXmlFile(file, resource, type_, itemSettings)
	}

	return nil
}

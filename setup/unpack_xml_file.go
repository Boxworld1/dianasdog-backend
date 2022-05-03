// @title   UnpackXmlFile
// @description 拆解 XML 并入库
// @auth	ryl			2022/4/26	13:30
// @param	filename	string		文件名
// @param	resource	string		特型卡类型
// @param	opType		string		操作类型（insert/delete）
// @param	itemSettings	[]getter.ItemSetting	写入行为
// @return  err			error		non-nil when fileName is wrong

package setup

import (
	"dianasdog/database"
	"dianasdog/getter"
)

func UnpackXmlFile(filename string, resource string, opType string, itemSettings []getter.ItemSetting) error {

	// 取得 xml 数据
	data, err := database.GetFile(database.DataClient, resource, filename)
	if err != nil {
		return err
	}

	UnpackXmlData(data, resource, opType, itemSettings, filename)

	return nil
}

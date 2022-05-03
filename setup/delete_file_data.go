// @title   DeleteFileData
// @description 删除单一文件
// @auth	ryl			2022/5/3	22:00
// @param	resource	string		特型卡类型
// @param	filename	string		文件名
// @return  err			error

package setup

import (
	"dianasdog/database"
)

func DeleteFileData(resource string, filename string) error {

	// 查找某一文件
	data, err := database.GetAllDocidByFilename(database.DocidClient, resource, filename)

	// 若特型卡类型错误
	if err != nil {
		return err
	}

	// 数据分解
	for _, block := range data {
		// 取得 docid
		docid := block.Name
		// 全局删除
		DeleteItem(resource, docid, 0)
	}

	return nil
}

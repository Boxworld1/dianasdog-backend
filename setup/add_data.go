// @title		AddData
// @description	新增数据接口
// @auth		ryl				2022/4/7		10:00
// @param		path			string			路径
// @param		targetResource	string			特型卡片类型（如 "诗词" 和 "车" 等）
// @return		err				error			错误值

package setup

func AddData(path string, targetResource string) error {

	// 临时以文件形式读入
	filesPath, err := GetXMLFiles(path)
	if err != nil {
		return err
	}

	// 遍历路径下所有 XML 文件
	for _, file := range filesPath {
		// 解码
		itemList, _, _, _, err := UnpackXMLFile(file, targetResource)
		if err != nil {
			return err
		}
		// 将每一条数据放入数据库中
		for _, item := range itemList {
			StoreItem(item, targetResource, "insert", "docid")
		}
	}

	return nil
}

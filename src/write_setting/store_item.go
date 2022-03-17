// @title		StoreItem
// @description	此函数会拆解上层接口提供的 item，并根据 get_config 的配置，将 item 中的数据截取并分类放入不同的数据库中。
// @auth		ryl				2022/3/17		11:05
// @param		data			*etree.Element	上层数据提供的 item 树
// @param		resourceName	string			特型卡类型
// @param		operation		string			对数据库的操作
// @return		err				error			错误值

package write_setting

import (
	"github.com/beevik/etree"
)

func StoreItem(data *etree.Element, resourceName string, operation string) error {
	var itemSettings []ItemSetting
	var err error

	// 查找对应特型卡的配置
	itemSettings, err = GetConfig(resourceName)
	if err != nil {
		return err
	}

	// 根据配置信息写入数据库
	for _, value := range itemSettings {

		// 根据路径选取对应数据
		// path := value.item_path

		// 数据写入摘要
		if value.dumpDigest {

		}

		// 数据写入倒排引擎
		if value.dumpInvertIdx {

		}

		// 数据写入词典
		if value.dumpDict {

		}
	}
	return nil
}

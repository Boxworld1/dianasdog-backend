// @title		StoreItem
// @description	此函数会拆解上层接口提供的 item，并根据 get_config 的配置，将 item 中的数据截取并分类放入不同的数据库中。
// @auth		ryl				2022/4/6		00:00
// @param		data			*etree.Element	上层数据提供的 item 树
// @param		resourceName	string			特型卡类型
// @param		operation		string			对数据库的操作
// @param		docid			string			item 编号
// @return		err				error			错误值

package setup

import (
	"dianasdog/database"
	"dianasdog/io"
	"strings"

	"github.com/beevik/etree"
)

func GetKey(resource string, tarPath string) string {
	path := strings.Split(tarPath, ".")
	key := path[len(path)-1]
	return resource + "@" + key
}

func StoreItem(data *etree.Element, resource string, operation string, docid string) error {
	var itemSettings []io.ItemSetting
	var err error

	// 查找对应特型卡的配置
	itemSettings, err = io.GetConfig(resource)
	if err != nil {
		return err
	}

	// 开启数据库
	redis := database.RedisClient
	es := database.EsClient
	dict := database.DictClient

	// 提取配置路径并在数据库中新增对应的表
	// for _, itemSetting := range itemSettings {
	// 	key := GetKey(resource, itemSetting.itemPath)
	// 	database.CreateTableFromDict(key)
	// }

	// 根据配置信息写入数据库
	for _, itemSetting := range itemSettings {

		// 根据路径选取对应数据
		// key := GetKey(resource, itemSetting.itemPath)
		path := strings.Replace(itemSetting.ItemPath, ".", "/", -1)

		for _, value := range data.FindElements(path) {
			// 数据写入摘要(Radis)
			if itemSetting.DumpDigest {
				// fmt.Println("insert to redis: ", value.Text())
				database.SetToRedis(redis, docid, value.Text())
			}

			// 数据写入倒排引擎(Es)
			if itemSetting.DumpInvertIdx {
				// fmt.Println("inesrt to es: ", value.Text())
				database.InsertToEs(es, docid, value.Text())
			}

			// 数据写入词典(Dict)
			if itemSetting.DumpDict {
				// fmt.Println("insert to dict", value.Text())
				database.InsertToDict(dict, resource, []string{docid, value.Text()})
			}
		}

	}
	return nil
}

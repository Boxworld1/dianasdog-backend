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
	"dianasdog/getter"
	"fmt"
	"strings"

	"github.com/beevik/etree"
)

func StoreItem(data *etree.Element, resource string, operation string, docid string, itemSettings []getter.ItemSetting) error {

	// 开启数据库
	redis := database.RedisClient
	es := database.EsClient

	// 新建表格
	database.CreateTableInDict(resource)
	database.CreateTableInPattern(resource)

	// 先记录要传入 Redis 的值
	var redisStr string
	var esStr string

	// 根据配置信息写入数据库
	for _, itemSetting := range itemSettings {

		// 根据路径选取对应数据
		key := GetKey(itemSetting.ItemPath)
		path := strings.Replace(itemSetting.ItemPath, ".", "/", -1)

		for _, value := range data.FindElements(path) {
			// 写入摘要(Radis)的数据
			if itemSetting.DumpDigest {
				// 若字串不为空，则加入逗号分隔
				if len(redisStr) != 0 {
					redisStr += ", "
				}
				// 然后存入数据
				redisStr += "\"" + key + "\": \"" + value.Text() + "\""
			}

			// 写入倒排引擎(Es)的数据
			if itemSetting.DumpInvertIdx {
				esStr = esStr + value.Text() + " "
			}

			// 数据写入词典(Dict)
			if itemSetting.DumpDict {
				fmt.Println("insert to dict", value.Text())
				database.InsertToDict(resource, docid, key, value.Text())
			}
		}

	}

	// 将传入 redis 的数据变为 json
	redisStr = "{" + redisStr + "}"

	// 数据写入摘要(Radis)
	fmt.Println("insert to redis: ", redisStr)
	database.SetToRedis(redis, docid, redisStr)

	// 数据写入倒排引擎(Es)
	fmt.Println("inesrt to es: ", esStr)
	database.InsertToEs(resource, es, docid, esStr)

	return nil
}

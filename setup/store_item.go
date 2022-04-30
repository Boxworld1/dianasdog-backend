// @title	StoreItem
// @description	此函数会拆解上层接口提供的 item，并根据 get_config 的配置，将 item 中的数据截取并分类放入不同的数据库中。
// @auth	ryl			2022/4/28		15:00
// @param	data		*etree.Element	上层数据提供的 item 树
// @param	resource	string			特型卡类型
// @param	docid		string			item 编号
// @return	err			error			错误值

package setup

import (
	"dianasdog/database"
	"dianasdog/getter"
	"fmt"
	"strconv"
	"strings"

	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/beevik/etree"
)

func StoreItem(data *etree.Element, resource string, docid string, itemSettings []getter.ItemSetting) error {

	// 开启数据库
	redis := database.RedisClient
	es := database.EsClient

	// 先初始化要传入 Redis 和 ES 的值
	var redisStr string
	var esStr string

	// 初始化 json
	myJson := jsonvalue.NewObject()
	myJson.SetString(resource).At("type")

	// 图片计数器
	picCount := 0

	// 数组计数器
	var itemCount map[string]int = make(map[string]int)

	// 根据配置信息写入数据库
	for _, itemSetting := range itemSettings {

		// 根据路径选取对应数据
		key := GetKey(itemSetting.ItemPath)
		keySlice := strings.Split(itemSetting.ItemPath, ".")
		path := strings.Replace(itemSetting.ItemPath, ".", "/", -1)

		// 查找所有可行路径
		for _, value := range data.FindElements(path) {

			// 索引初始化
			var pathKey string = ""

			// 将 []string 拆为 []interface{}
			pathList := make([]interface{}, 0)
			for _, key1 := range keySlice {
				// 索引拼接
				pathKey = pathKey + "@" + key1

				// 路径加长
				pathList = append(pathList, key1)

				// 若为特定键值
				if key1 == "item" || key1 == "tag" || key1 == "tab" {
					// 检查键值是否出现过
					if _, ok := itemCount[pathKey]; !ok {
						itemCount[pathKey] = 0
					}
					// 设置位置
					pathList = append(pathList, itemCount[pathKey])
					itemCount[pathKey]++

					// 更改位置
					pathKey = pathKey + "@" + strconv.Itoa(itemCount[pathKey]-1)
				}
			}

			// 写入摘要(Redis)的数据
			if itemSetting.DumpDigest {
				// 若为图片
				if itemSetting.IsPic {
					myJson.SetString(value.Text()).At("picture", picCount)
					picCount++
				} else {
					// 然后插入 Json
					myJson.SetString(value.Text()).At("item", pathList...)
				}
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
	redisStr = myJson.MustMarshalString()

	// 数据写入摘要(Radis)
	fmt.Println("insert to redis: ", redisStr)
	database.SetToRedis(redis, docid, redisStr)

	// 数据写入倒排引擎(Es)
	fmt.Println("inesrt to es: ", esStr)
	database.InsertToEs(resource, es, docid, esStr)

	return nil
}

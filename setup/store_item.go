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
	"strings"

	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/beevik/etree"
)

func isSpecial(key string) bool {
	if key == "item" || key == "tag" || key == "tab" || key == "chapter_info" {
		return true
	}
	return false
}

func dfs(data *etree.Element, keySlice []string, path []interface{}, res string, docid string,
	itemSetting getter.ItemSetting, myJson *jsonvalue.V, picCount *int) {

	// Json Tree 索引记录
	pathList := path
	// 初始化 etree 索引路径
	var nowPath string = data.GetPath()

	for idx, keyValue := range keySlice {
		// 路径加长
		pathList = append(pathList, keyValue)
		// 索引路径，先加分隔线然后接上当前路径
		nowPath += "/" + keyValue

		// 若到达目标位置
		if idx == len(keySlice)-1 {
			for _, value := range data.FindElements(nowPath) {
				// 写入摘要(Redis)的数据
				if itemSetting.DumpDigest {
					// 若为图片
					if itemSetting.IsPic {
						myJson.SetString(value.Text()).At("picture", *picCount)
						*picCount++
					} else {
						// 然后插入 Json
						myJson.SetString(value.Text()).At("item", pathList...)
					}
				}
			}
			break
		}

		// 若为特定键值
		if isSpecial(keyValue) {
			for cnt, value := range data.FindElements(nowPath) {
				// 设置元素在 Json Tree 的位置
				tmpPath := append(pathList, cnt)
				doc := etree.NewDocument()
				doc.SetRoot(value.Copy())
				dfs(doc.Root(), keySlice[idx+1:], tmpPath, res, docid, itemSetting, myJson, picCount)
			}
			break
		}
	}
}

func StoreItem(data *etree.Element, resource string, docid string, itemSettings []getter.ItemSetting) error {

	// 开启数据库
	redis := database.RedisClient
	es := database.EsClient

	// 先初始化要传入 Redis 和 ES 的值
	redisStr := ""
	esStr := ""

	// 图片计数器
	picCount := 0

	// 初始化 json
	myJson := jsonvalue.NewObject()
	myJson.SetString(resource).At("type")

	// 根据配置信息写入数据库
	for _, itemSetting := range itemSettings {

		// 根据路径选取对应数据
		keySlice := strings.Split(itemSetting.ItemPath, ".")
		path := strings.Replace(itemSetting.ItemPath, ".", "/", -1)
		var pathList []interface{} = make([]interface{}, 0)

		// 递归查找(Redis)
		dfs(data, keySlice, pathList, resource, docid, itemSetting, myJson, &picCount)

		for _, value := range data.FindElements(path) {
			// 写入倒排引擎(Es)的数据
			if itemSetting.DumpInvertIdx {
				esStr = esStr + value.Text() + " "
			}

			// 数据写入词典(Dict)
			if itemSetting.DumpDict {
				fmt.Println("insert to dict", value.Text())
				database.InsertToDict(resource, docid, keySlice[len(keySlice)-1], value.Text())
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

// @title		GetConfig
// @description	此函数的用途为，根据数据 "类型"，在配置文件中找出对应的 "写入行为"，并反馈相关数据到数据处理函数中。
// @auth		ryl				2022/3/16		21:00
// @param		target_resource	string			特型卡片类型（如 "诗词" 和 "车" 等）
// @param		target_key		string			欲查找的键值
// @returns		ItemSettings	[]ItemSetting	此键值下所有需要写入数据库的数据
// @return		err				error			错误值

package write_setting

import (
	"io/ioutil"
	"strings"

	"github.com/tidwall/gjson"
)

// ItemSetting	保存 存入数据库的数据在文件中的路径 和 需要储入的数据库
type ItemSetting struct {
	item_path       string // 存入数据库的资料路径
	dump_digest     bool   // 本字段是否需要 dump 摘要
	dump_invert_idx bool   // 本字段是否需要 dump 倒排
	dump_dict       bool   // 本字段是否需要 dump 词表
}

func GetConfig(target_resource string, target_key string) ([]ItemSetting, error) {

	// TODO	未来可以利用 map 查找对应类型的 config 文档路径
	filepath := "./testcase/test.json"
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	json := string(file)
	settings := gjson.Get(json, "write_setting|@pretty")

	var ItemSettings = make([]ItemSetting, 0)

	// TODO 在配置文件中查找可能的配置
	settings.ForEach(func(key, value gjson.Result) bool {

		// 先查找前缀相同的数据，这是因为键值中不含"."
		if strings.HasPrefix(key.String(), target_key+".") {
			var item ItemSetting
			item.item_path = key.String()

			// 读取此路径下的 dump 信息
			value.ForEach(func(key, value gjson.Result) bool {
				switch key.String() {
				case "dump_digest":
					item.dump_digest = value.Bool()
				case "dump_dict":
					item.dump_dict = value.Bool()
				case "dump_invert_idx":
					item.dump_invert_idx = value.Bool()
				}
				return true
			})

			// 更新配置数组
			ItemSettings = append(ItemSettings, item)
		}
		return true
	})

	// TODO 查找成功，返回数组
	return ItemSettings, nil
}

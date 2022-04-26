// @title	SetConfig
// @description	此函数的用途为，根据数据"类型"写入配置文件，然后根据新旧配置文件的对比结果全量建库。
// @auth	ryl			2022/4/26		16:00
// @param	resource	string			特型卡片类型（如 "诗词" 和 "车" 等）
// @param	content		[]ItemSetting	需要写入配置文件的数据
// @return	err			error			错误值

package setter

import (
	"dianasdog/database"
	"dianasdog/getter"
	"dianasdog/setup"
)

func SetConfig(resource string, content []byte) error {

	// 查找原来特型卡的配置
	oldSetting, _ := getter.GetConfig(resource)

	// 写入新配置
	err := database.InsertFile(database.ConfigClient, "file", resource, content)
	if err != nil {
		return err
	}

	// 查找新特型卡配置
	newSetting, err := getter.GetConfig(resource)

	// 删除信息 0:删除, 1:维持/不操作
	var oldStatus []int = make([]int, len(oldSetting))
	// 更改信息 0:更新/插入, 1:维持/不操作
	var newStatus []int = make([]int, len(newSetting))

	// 对比两文件之差
	for newID, newItem := range newSetting {
		for oldID, oldItem := range oldSetting {
			// 查找相同路径的配置
			if newItem.ItemPath == oldItem.ItemPath {
				// 旧信息不需操作
				oldStatus[oldID] = 1
				// 检查信息，若配置不变
				if (newItem.DumpDict == oldItem.DumpDict) || (newItem.DumpDigest == oldItem.DumpDigest) || (newItem.DumpInvertIdx == oldItem.DumpInvertIdx) {
					// 维持即可
					newStatus[newID] = 1
				}
			}
		}
	}

	// 更改对应文件入库的数据（全量建库）
	var deleteSetting []getter.ItemSetting = make([]getter.ItemSetting, 0)
	var insertSetting []getter.ItemSetting = make([]getter.ItemSetting, 0)

	// 统计删除信息
	for oldID, oldItem := range oldSetting {
		// 若需要删除
		if oldStatus[oldID] == 0 {
			deleteSetting = append(deleteSetting, oldItem)
		}
	}

	// 统计插入信息
	for newID, newItem := range oldSetting {
		// 若需要删除
		if newStatus[newID] == 0 {
			insertSetting = append(insertSetting, newItem)
		}
	}

	// 先删除数据
	go setup.UpdateResData(resource, "delete", deleteSetting)

	// 再更新数据
	go setup.UpdateResData(resource, "insert", insertSetting)

	// 无论正确与否都返回 err 的内容
	return err
}

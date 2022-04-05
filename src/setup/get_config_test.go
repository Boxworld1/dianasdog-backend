// @title		TestGetConfig
// @description	此函数的用途为检查 GetConfig 函数的正确性
// @auth		ryl				2022/3/17		10:00
// @param		t				*testing.T		testing 用参数

package setup

import (
	"testing"
)

func TestGetConfig(t *testing.T) {

	// 查找不存在的文件
	_, err := GetConfig("apple")
	if err == nil {
		t.Error(err)
	}

	itemSettings, err := GetConfig("car")
	// 测试时出错
	if err != nil {
		t.Error(err)
	}

	// 配置数量错误
	if len(itemSettings) != 4 {
		t.Error("配置数量有误")
	} else {
		item0 := itemSettings[0]
		item1 := itemSettings[1]
		// 配置错误
		if item0.dumpDigest != true || item1.dumpDigest != false {
			t.Error("dump digest 错误")
		}
		if item0.dumpInvertIdx != false || item1.dumpInvertIdx != false {
			t.Error("dump invert idx 错误")
		}
		if item0.dumpDict != true || item1.dumpDict != true {
			t.Error("dump dict 错误")
		}

	}
}

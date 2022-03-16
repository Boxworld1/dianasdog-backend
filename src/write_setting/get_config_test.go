// @title		TestGetConfig
// @description	此函数的用途为检查 GetConfig 函数的正确性
// @auth		ryl				2022/3/16		22:30
// @param		t				*testing.T		testing 用参数

package write_setting

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	ItemSettings, err := GetConfig("car", "a")
	// 测试时出错
	if err != nil {
		t.Error(err)
	}
	// 配置数量错误
	if len(ItemSettings) != 2 {
		t.Error("配置数量有误")
	} else {
		item0 := ItemSettings[0]
		item1 := ItemSettings[1]
		// 配置错误
		if item0.dump_digest != true || item1.dump_digest != false {
			t.Error("dump digest 错误")
		}
		if item0.dump_invert_idx != false || item1.dump_invert_idx != false {
			t.Error("dump invert idx 错误")
		}
		if item0.dump_dict != true || item1.dump_dict != true {
			t.Error("dump dict 错误")
		}

	}
}

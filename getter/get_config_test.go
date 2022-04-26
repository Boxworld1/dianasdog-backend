// @title		TestGetConfig
// @description	此函数的用途为检查 GetConfig 函数的正确性
// @auth		ryl				2022/3/17		10:00
// @param		t				*testing.T		testing 用参数

package getter

import (
	"dianasdog/testcase"
	"testing"
)

func TestGetConfig(t *testing.T) {

	// 初始化测例
	if err := testcase.SetTestData(0); err != nil {
		t.Error("测例建造失败")
	}

	// 查找不存在的文件
	_, err := GetConfig("testcase_apple")
	if err == nil {
		t.Error(err)
	}

	itemSettings, err := GetConfig("testdata")
	// 测试时出错
	if err != nil {
		t.Error(err)
	}

	// 定义结果
	ans := [][]bool{{true, true, true}, {true, false, true}}

	// 配置数量错误
	if len(itemSettings) != 2 {
		t.Error("配置数量有误", len(itemSettings))
	} else {
		for _, item := range itemSettings {
			// 配置错误
			key := 0
			if item.ItemPath == "display.title" {
				key = 1
			}
			if item.DumpDigest != ans[key][0] {
				t.Error("dump digest 错误")
			}
			if item.DumpInvertIdx != ans[key][1] {
				t.Error("dump invert idx 错误")
			}
			if item.DumpDict != ans[key][2] {
				t.Error("dump dict 错误")
			}
		}

	}
}

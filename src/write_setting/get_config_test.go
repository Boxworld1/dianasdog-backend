package write_setting

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	ItemSettings, err := GetConfig("car", "a")
	if err != nil {
		t.Error(err)
	}
	if len(ItemSettings) != 1 {
		t.Error("配置有误")
	} else {
		item := ItemSettings[0]
		if item.dump_dict != true {
			t.Error("dump dict 错误")
		}
		if item.dump_invert_idx != false {
			t.Error("dump invert idx 错误")
		}
		if item.dump_digest != true {
			t.Error("dump digest 错误")
		}
	}
}

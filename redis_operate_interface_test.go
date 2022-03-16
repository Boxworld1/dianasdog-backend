package matianyu

import "testing"

func TestRedisOperationInterface(t *testing.T) {
	var client = connect_to_redis()
	if client == nil {
		t.Error("connect_to_redis function get a nil pointer for client")
	} else if set_to_redis(client, "name", "matianyu") {
		value, succ := get_from_redis(client, "name")
		if !succ {
			t.Error("get_from_redis function generates error")
		} else {
			if value != "matianyu" {
				t.Errorf("Expected value is matianyu, but %s got", value)
			}
		}
		if delete_from_redis(client, "name") {
			exist, succ := exist_in_redis(client, "name")
			if !succ {
				t.Error("exist_in_redis function generates error")
			} else {
				if exist {
					t.Errorf("delete_from_redis function don't delete asked element")
				}
			}
		} else {
			t.Error("delete_from_redis generates error")
		}
	} else {
		t.Error("set_to_redis function generates error")
	}

}

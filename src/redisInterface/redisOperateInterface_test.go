package matianyu

import "testing"

func TestRedisOperationInterface(t *testing.T) {
	var client = connectToRedis()
	if client == nil {
		t.Error("connectToRedis function get a nil pointer for client")
	} else if setToRedis(client, "name", "matianyu") {
		value, succ := getFromRedis(client, "name")
		if !succ {
			t.Error("getFromRedis function generates error")
		} else {
			if value != "matianyu" {
				t.Errorf("Expected value is matianyu, but %s got", value)
			}
		}
		if deleteFromRedis(client, "name") {
			exist, succ := existInRedis(client, "name")
			if !succ {
				t.Error("existInRedis function generates error")
			} else {
				if exist {
					t.Errorf("deleteFromRedis function don't delete asked element")
				}
			}
		} else {
			t.Error("deleteFromRedis generates error")
		}
	} else {
		t.Error("setToRedis function generates error")
	}

}

// @title	TestRedisOperationInterface
// @description	此函数的用途为检查 redis 的接口函数正确性
// @auth	mdy		2022/3
// @param	t		*testing.T	testing 用参数

package database

import "testing"

func TestRedisOperationInterface(t *testing.T) {
	var client = RedisClient
	if client == nil {
		t.Error("connectToRedis function get a nil pointer for client")
	} else if SetToRedis(client, "name", "matianyu") {
		value, succ := GetFromRedis(client, "name")
		if !succ {
			t.Error("getFromRedis function generates error")
		} else {
			if value != "matianyu" {
				t.Errorf("Expected value is matianyu, but %s got", value)
			}
		}
		if DeleteFromRedis(client, "name") {
			exist, succ := ExistInRedis(client, "name")
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

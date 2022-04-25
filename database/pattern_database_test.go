package database

import (
	"testing"
)

func TestPatternInterface(t *testing.T) {
	//测试创建表
	err := CreateTableInPattern("testcase")
	if err != nil {
		t.Error(err)
	}

	//测试插入数据
	err = InsertToPattern("testcase", "title+garbage+intent+garbage")
	if err != nil {
		t.Error(err)
	}

	//测试读取数据
	res, err := FetchAllPattern("testcase")
	if res[0] != "title+garbage+intent+garbage" || err != nil {
		t.Error("Fail to fetch patterns")
	}

	//测试删除数据
	err = DeleteFromPattern("testcase", "title+garbage+intent+garbage")
	if err != nil {
		t.Error(err)
	}

	//测试删除表
	err = DeleteTableFromPattern("testcase")
	if err != nil {
		t.Error(err)
	}
}

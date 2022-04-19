// @title	TestDictInterface
// @description	检查 sql 的功能
// @auth	jz		2022/3/30
// @auth	ryl		2022/4/19	17:30
// @param	t		*testing.T

package database

import (
	"testing"
)

func TestDictInterface(t *testing.T) {
	// 创建数据库 dict
	db, err := CreateDatabase("dict")
	if err != nil {
		t.Error(err)
	}

	// 新建表格
	err = CreateTableFromDict(db, "car")
	if err != nil {
		t.Error(err)
	}

	// 插入数据
	err = InsertToDict(db, "car", "奔驰")
	if err != nil {
		t.Error(err)
	}
	err = InsertToDict(db, "car", "奔驰")
	if err != nil {
		t.Error(err)
	}
	err = InsertToDict(db, "flower", "宝马")
	if err == nil {
		t.Error("禁止向不存在的表中插入数据")
	}

	// 查找数据
	tmp, _ := SearchFromDict(db, "car", "flower")
	if tmp == "flower" {
		t.Error("flower is not in car")
	}
	tmp, err = SearchFromDict(db, "car", "奔驰")
	if tmp != "奔驰" {
		t.Error("查询失败")
	}
	if err != nil {
		t.Error(err)
	}

	// 删除数据
	err = DeleteFromDict(db, "car", "奔驰")
	if err != nil {
		t.Error(err)
	}
	err = DeleteFromDict(db, "flower", "宝马")
	if err == nil {
		t.Error("禁止向不存在的表中删除数据")
	}

	// 删除表格
	err = DeleteTableFromDict(db, "car")
	if err != nil {
		t.Error(err)
	}
}

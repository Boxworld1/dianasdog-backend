package database

import (
	"testing"
)

// test for function: createTable
func TestCreateTableFromDict(t *testing.T) {
	err := CreateTableFromDict("car")
	if err != nil {
		t.Error(err)
	}
}

// test for function: insert
func TestInsertToDict(t *testing.T) {
	err := InsertToDict("car", "奔驰")
	if err != nil {
		t.Error(err)
	}
	err = InsertToDict("car", "奔驰")
	if err != nil {
		t.Error(err)
	}
	err = InsertToDict("flower", "宝马")
	if err == nil {
		t.Error("禁止向不存在的表中插入数据")
	}
}

// test for function: search
func TestSearchFromDict(t *testing.T) {
	tmp, _ := SearchFromDict("car", "flower")
	if tmp == "flower" {
		t.Error("flower is not in car")
	}
	tmp, err := SearchFromDict("car", "奔驰")
	if tmp != "奔驰" {
		t.Error("查询失败")
	}
	if err != nil {
		t.Error(err)
	}

}

// test for function: delete
func TestDeleteFromDict(t *testing.T) {
	err := DeleteFromDict("car", "奔驰")
	if err != nil {
		t.Error(err)
	}
	err = DeleteFromDict("flower", "宝马")
	if err == nil {
		t.Error("禁止向不存在的表中删除数据")
	}
}

// test for function: deleteTable
func TestDeleteTableFromDict(t *testing.T) {
	err := DeleteTableFromDict("car")
	if err != nil {
		t.Error(err)
	}
}

package dict

import (
	"testing"
)

// test for function: createTable
func TestCreateTable(t *testing.T) {
	err := CreateTable("car")
	if err != nil {
		t.Error(err)
	}
}

// test for function: insert
func TestInsert(t *testing.T) {
	err := Insert("car", "奔驰")
	if err != nil {
		t.Error(err)
	}
	err = Insert("car", "奔驰")
	if err != nil {
		t.Error(err)
	}
	err = Insert("flower", "宝马")
	if err == nil {
		t.Error("禁止向不存在的表中插入数据")
	}
}

// test for function: search
func TestSearch(t *testing.T) {
	tmp, _ := Search("car", "flower")
	if tmp == "flower" {
		t.Error("flower is not in car")
	}
	tmp, err := Search("car", "奔驰")
	if tmp != "奔驰" {
		t.Error("查询失败")
	}
	if err != nil {
		t.Error(err)
	}

}

// test for function: delete
func TestDelete(t *testing.T) {
	err := Delete("car", "奔驰")
	if err != nil {
		t.Error(err)
	}
	err = Delete("flower", "宝马")
	if err == nil {
		t.Error("禁止向不存在的表中删除数据")
	}
}

// test for function: deleteTable
func TestDeleteTable(t *testing.T) {
	err := DeleteTable("car")
	if err != nil {
		t.Error(err)
	}
}

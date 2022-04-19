package database

import (
	"testing"
)

// test for function: createTable
func TestCreateTableFromDict(t *testing.T) {
	d := []string{"id", "title"}
	err := CreateTableFromDict("car", d)
	if err != nil {
		t.Error(err)
	}
}

//test for function: showtablesindict
func TestShowTablesInDict(t *testing.T) {
	dict, err := ShowTablesInDict()
	if err != nil {
		t.Error(err)
	}
	var exist = false
	for _, tmp := range dict {
		if tmp == "car" {
			exist = true
			break
		}
	}
	if !exist {
		t.Error("返回错误")
	}
}

//test for function:ShowColumnsInTable
func TestShowColumnsInTable(t *testing.T) {
	dict, err := ShowColumnsInTable("car")
	if err != nil {
		t.Error(err)
	}
	if dict[0] != "id" || dict[1] != "title" {
		t.Error("返回错误")
	}
}

// test for function: insert
func TestInsertToDict(t *testing.T) {
	item := []string{"test0", "奔驰"}
	err := InsertToDict("car", item)
	if err != nil {
		t.Error(err)
	}
	err = InsertToDict("car", item)
	if err != nil {
		t.Error(err)
	}
	item = []string{"test1", "宝马"}
	err = InsertToDict("flower", item)
	if err == nil {
		t.Error("禁止向不存在的表中插入数据")
	}
}

// test for function: search
func TestSearchFromDict(t *testing.T) {
	tmp, _ := SearchFromDict("car", "test2")
	if tmp[0] == "test2" {
		t.Error("test2 is not in car")
	}
	tmp, err := SearchFromDict("car", "test0")
	if tmp[0] != "test0" || tmp[1] != "奔驰" {
		t.Error("查询失败")
	}
	if err != nil {
		t.Error(err)
	}

}

//test for function QueryColumn
func TestQueryColumn(t *testing.T) {
	item := []string{"test1", "宝马"}
	InsertToDict("car", item)
	dictionary, err := QueryColumn("car", "title")
	if err != nil {
		t.Error(err)
	}
	check1 := (dictionary[0] == "奔驰" && dictionary[1] == "宝马")
	check2 := (dictionary[0] == "宝马" && dictionary[1] == "奔驰")
	if !check1 && !check2 {
		t.Error("返回错误")
	}

}

// test for function: delete
func TestDeleteFromDict(t *testing.T) {
	err := DeleteFromDict("car", "test1")
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

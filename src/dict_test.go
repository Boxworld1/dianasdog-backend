// @Title  dict
// @Description  construct the local dictionary and support the insert, search and delete operations
// @Author  蒋政
// @Update  2022/3/16
package dict

import (
	"testing"
)

// test for function: create_table
func TestCreateTable(t *testing.T) {
	err := CreateTable("car")
	if err != nil {
		t.Error(err)
	}
}

// test for function: insert
func TestInsert(t *testing.T) {
	err := Insert("奔驰", "car")
	if err != nil {
		t.Error(err)
	}
	err = Insert("奔驰", "car")
	if err != nil {
		t.Error(err)
	}
	err = Insert("宝马", "flower")
	if err == nil {
		t.Error("禁止向不存在的表中插入数据")
	}
}

// test for function: search
func TestSearch(t *testing.T) {
	tmp, _ := Search("flower", "car")
	if tmp == "flower" {
		t.Error("flower is not in car")
	}
	tmp, err := Search("奔驰", "car")
	if tmp != "奔驰" {
		t.Error("返回错误")
	}
	if err != nil {
		t.Error(err)
	}

}

// test for function: delete
func TestDelete(t *testing.T) {
	err := Delete("奔驰", "car")
	if err != nil {
		t.Error(err)
	}
	err = Delete("宝马", "flower")
	if err == nil {
		t.Error("禁止向不存在的表中删除数据")
	}
}

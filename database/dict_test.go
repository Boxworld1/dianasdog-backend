// @title	TestDbInterface
// @description	此函数的用途为检查 sql 的接口函数正确性
// @auth	jz		2022/4/19	23:30
// @auth	ryl		2022/4/20	10:30
// @param	t		*testing.T	testing 用参数

package database

import (
	"testing"
)

func TestDbInterface(t *testing.T) {
	// 取得数据库
	db := DictClient

	// 新建表格
	d := []string{"id", "title"}
	err := CreateTableFromDict(db, "testcase", d)
	if err != nil {
		t.Error(err)
	}

	// 测试表格是否正确插入
	dict, err := ShowTablesInDict(db)
	if err != nil {
		t.Error(err)
	}
	var exist = false
	for _, tmp := range dict {
		if tmp == "testcase" {
			exist = true
			break
		}
	}
	if !exist {
		t.Error("返回错误")
	}

	// 检查列名是否正确生成
	dict, err = ShowColumnsInTable(db, "testcase")
	if err != nil {
		t.Error(err)
	}
	if dict[0] != "id" || dict[1] != "title" {
		t.Error("返回错误")
	}

	// 测试插入功能：向已存在的表中插入数据
	item := []string{"test0", "奔驰"}
	err = InsertToDict(db, "testcase", item)
	if err != nil {
		t.Error(err)
	}
	// 测试插入功能：重复插入数据
	err = InsertToDict(db, "testcase", item)
	if err != nil {
		t.Error(err)
	}

	// 测试插入功能：向不存在的表中插入数据
	item = []string{"test1", "宝马"}
	err = InsertToDict(db, "testcase_flower", item)
	if err == nil {
		t.Error("禁止向不存在的表中插入数据")
	}

	// 测试搜索功能：查找不存在的键值
	tmp, _ := SearchFromDict(db, "testcase", "test2")
	if len(tmp) != 0 {
		t.Error("test2 is not in testcase")
	}

	// 测试搜索功能：查找存在的键值
	tmp, err = SearchFromDict(db, "testcase", "test0")
	if tmp[0] != "test0" || tmp[1] != "奔驰" {
		t.Error("查询失败")
	}
	if err != nil {
		t.Error(err)
	}

	// 按列查询所有元素
	item = []string{"test1", "宝马"}
	InsertToDict(db, "testcase", item)
	dictionary, err := QueryColumn(db, "testcase", "title")
	if err != nil {
		t.Error(err)
	}
	check1 := (dictionary[0] == "奔驰" && dictionary[1] == "宝马")
	check2 := (dictionary[0] == "宝马" && dictionary[1] == "奔驰")
	if !check1 && !check2 {
		t.Error("返回错误")
	}

	// 测试删除功能：向存在的表中作删除
	err = DeleteFromDict(db, "testcase", "test1")
	if err != nil {
		t.Error(err)
	}

	// 测试删除功能：向不存在的表中作删除
	err = DeleteFromDict(db, "testcase_flower", "宝马")
	if err == nil {
		t.Error("禁止向不存在的表中删除数据")
	}

	// 测试表格删除功能
	err = DeleteTableFromDict(db, "testcase")
	if err != nil {
		t.Error(err)
	}
}

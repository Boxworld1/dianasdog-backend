// @title	TestDbInterface
// @description	此函数的用途为检查 sql 的接口函数正确性
// @auth	jz		2022/4/25	11:27
// @auth	ryl		2022/4/20	10:30
// @param	t		*testing.T	testing 用参数

package database

import (
	"testing"
)

func TestDbInterface(t *testing.T) {
	// 新建表格
	err := CreateTableInDict("testcase")
	if err != nil {
		t.Error(err)
	}

	// 测试表格是否正确插入
	dict, err := ShowTablesInDict()
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
		t.Error("建表失败")
	}

	// 测试插入功能：向已存在的表中插入数据
	err = InsertToDict("testcase", "10086", "title", "奔驰")
	if err != nil {
		t.Error(err)
	}
	// 测试插入功能：重复插入数据
	err = InsertToDict("testcase", "10086", "title", "奔驰")
	if err != nil {
		t.Error(err)
	}

	// 测试插入功能：向不存在的表中插入数据
	err = InsertToDict("testcase_flower", "10086", "title", "奔驰")
	if err == nil {
		t.Error("禁止向不存在的表中插入数据")
	}

	// 测试搜索功能(SearchByDocid)：查找不存在的docid
	tmp, err := SearchByDocid("testcase", "12345")
	if len(tmp) != 0 {
		t.Error("Docid 12345 is not in testcase")
	}
	if err != nil {
		t.Error(err)
	}

	// 测试搜索功能(SearchByDocid)：查找存在的docid
	tmp, err = SearchByDocid("testcase", "10086")
	if tmp[0][0] != "title" || tmp[0][1] != "奔驰" {
		t.Error("查询失败")
	}
	if err != nil {
		t.Error(err)
	}

	// 测试搜索功能(SearchByField)
	tmp1, err := SearchByField("testcase", "10086", "title")
	if tmp1[0] != "奔驰" || err != nil {
		t.Error("查询失败")
	}

	//测试取词功能(GetAllWord):不存在的键值
	res, err := GetAllWord("testcase", "name")
	if len(res) != 0 {
		t.Error("Field name is not in testcase")
	}
	if err != nil {
		t.Error(err)
	}

	// 测试取词功能(GetAllWord)：存在的键值
	res, err = GetAllWord("testcase", "title")
	if res[0] != "奔驰" {
		t.Error("查询失败")
	}
	if err != nil {
		t.Error(err)
	}

	// 测试获取所有字段名
	dict, err = GetAllField("testcase")
	if len(dict) == 0 || err != nil || dict[0] != "title" {
		t.Error("GetAllField failed.")
	}

	//测试删除功能：删除某一field
	err = DeleteByField("testcase", "10086", "title")
	if err != nil {
		t.Error(err)
	}

	// 测试删除功能：向存在的表中作删除
	InsertToDict("testcase", "10086", "title", "奔驰")
	err = DeleteByDocid("testcase", "10086")
	if err != nil {
		t.Error(err)
	}

	// 测试删除功能：向不存在的表中作删除
	err = DeleteByDocid("testcase_flower", "10086")
	if err == nil {
		t.Error("禁止向不存在的表中删除数据")
	}

	// 测试表格删除功能
	err = DeleteTableFromDict("testcase")
	if err != nil {
		t.Error(err)
	}
}

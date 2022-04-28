// @title	TestCategoryTable
// @description	此函数的用途为检查 sql category 的接口函数正确性
// @auth	ryl		2022/4/27	21:30
// @param	t		*testing.T	testing 用参数

package database

import (
	"testing"
)

func TestCategoryTable(t *testing.T) {

	// 新增表格
	if err := CreateCategoryTable(CategoryClient, "hi"); err != nil {
		t.Error(err)
	}

	// 插入类別
	if err := InsertCategory(CategoryClient, "hi", "bye"); err != nil {
		t.Error(err)
	}

	// 取出所有类別
	if _, err := GetAllCategory(CategoryClient, "hi"); err != nil {
		t.Error(err)
	}

	// 取出所有类別（不存在）
	if _, err := GetAllCategory(CategoryClient, "hi23150"); err == nil {
		t.Error("不能检测错误")
	}

	// 统计存在的类別
	if _, err := CountCategory(CategoryClient, "hi"); err != nil {
		t.Error(err)
	}

	// 统计不存在的类別
	if _, err := CountCategory(CategoryClient, "hi34"); err == nil {
		t.Error("不能检测错误")
	}

	// 删除表格
	if err := DropCategory(CategoryClient, "hi"); err != nil {
		t.Error(err)
	}

	// 删除不存在的表格
	if err := DropCategory(CategoryClient, "bye"); err == nil {
		t.Error("不能检测错误")
	}

}

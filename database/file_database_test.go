// @title	TestFileTable
// @description	此函数的用途为检查 sql file 的接口函数正确性
// @auth	ryl		2022/4/20	14:30
// @param	t		*testing.T	testing 用参数

package database

import (
	"bytes"
	"testing"
)

func TestFileTable(t *testing.T) {
	// 设置文件名字及内容
	filename := "test.txt"
	srcData := []byte("my data")

	// 插入文件内容
	err := InsertFile(ConfigClient, "file", filename, srcData)
	if err != nil {
		t.Error(err)
	}

	// 取出文件内容
	dstData, err := GetFile(ConfigClient, "file", filename)
	if err != nil {
		t.Error(err)
	}

	// 对比结果
	res := bytes.Compare(srcData, dstData)
	if res != 0 {
		t.Error("存取错误!")
	}

	// 取出文件名
	_, err = GetFileName(ConfigClient, "file")
	if err != nil {
		t.Error("存取错误!")
	}

	// 取出不存在的文件
	_, err = GetFile(ConfigClient, "file", "filename")
	if err == nil {
		t.Error("检测不到错误")
	}

	// 取出不存在的表格
	_, err = GetFile(ConfigClient, "apple", "filename")
	if err == nil {
		t.Error("检测不到错误")
	}

}

// @title	TestFileTable
// @description	此函数的用途为检查 sql file 的接口函数正确性
// @auth	ryl		2022/4/20	14:30
// @param	t		*testing.T	testing 用参数

package database

import (
	"bytes"
	"testing"
)

func TestDocidTable(t *testing.T) {
	// 设置文件名字及内容
	filename := "test.txt"
	docid := "file@mytest"
	resource := "file"
	srcData := []byte("my data")
	CreateDocidTable(DocidClient, resource)

	// 插入 docid 内容
	err := InsertDocid(DocidClient, resource, docid, srcData, filename)
	if err != nil {
		t.Error(err)
	}

	// 取出 docid 内容
	dstData, err := GetDocid(DocidClient, resource, docid)
	if err != nil {
		t.Error(err)
	}

	// 对比结果
	res := bytes.Compare(srcData, dstData)
	if res != 0 {
		t.Error("存取错误!")
	}

	// 取出不存在的 docid
	_, err = GetDocid(DocidClient, resource, "mydocid")
	if err == nil {
		t.Error("检测不到错误")
	}

	// 取出不存在的表格
	_, err = GetDocid(DocidClient, "tes_apple", docid)
	if err == nil {
		t.Error("检测不到错误")
	}

	// 取出所有 docid
	_, err = GetAllDocid(DocidClient, resource)
	if err != nil {
		t.Error(err)
	}

	// 取出不存在的表格
	_, err = GetAllDocid(DocidClient, "3109275jk")
	if err == nil {
		t.Error("检测不到错误")
	}

	// 取出所有 docid
	_, err = GetAllDocidByFilename(DocidClient, resource, filename)
	if err != nil {
		t.Error(err)
	}

	// 取出不存在的表格
	_, err = GetAllDocidByFilename(DocidClient, "3109275jk", filename)
	if err == nil {
		t.Error("检测不到错误")
	}

	// 删除文件
	err = DeleteDocid(DocidClient, resource, docid)
	if err != nil {
		t.Error(err)
	}

	// 删除不存在的文件
	err = DeleteDocid(DocidClient, resource, "234979832")
	if err != nil {
		t.Error(err)
	}

	// 删除不存在的类型
	err = DeleteDocid(DocidClient, "1328104809328509", "234979832")
	if err == nil {
		t.Error("检测不到错误")
	}
}

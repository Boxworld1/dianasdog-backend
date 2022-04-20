// @title	TestSetTemplate
// @description	此函数的用途为检查 SetTemplate 函数的正确性
// @auth	ryl		2022/4/13	10:00
// @param	t		*testing.T	testing 用参数

package io

import (
	"dianasdog/path"
	"io/ioutil"
	"testing"
)

func TestSetTemplate(t *testing.T) {

	// 得到此文件的绝对路径
	abspath, _ := path.GetAbsPath()

	// 取得测试文件
	filepath := abspath + "template/testcase.json"

	// 读入文件
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Error(err)
	}

	err = SetTemplate("testcase", file)
	// 测试时出错
	if err != nil {
		t.Error(err)
	}

}

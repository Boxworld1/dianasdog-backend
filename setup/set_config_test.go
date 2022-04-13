// @title	TestSetConfig
// @description	此函数的用途为检查 SetConfig 函数的正确性
// @auth	ryl		2022/4/13	10:00
// @param	t		*testing.T	testing 用参数

package setup

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestSetConfig(t *testing.T) {

	// 得到此文件的绝对路径
	abspath, _ := GetAbsPath()

	// 取得测试文件
	filepath := abspath + "config/testcase_car.json"
	fmt.Println(filepath)
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Error(err)
	}

	err = SetConfig("testcase_car", file)
	// 测试时出错
	if err != nil {
		t.Error(err)
	}

}

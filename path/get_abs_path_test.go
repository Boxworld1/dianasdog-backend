// @title		TestGetAbsPath
// @description	检查 GetAbsPath 返回的项目绝对路径
// @auth		ryl				2022/4/7		12:00
// @param		t				*testing.T		testing 用参数

package path

import (
	"fmt"
	"testing"
)

func TestGetAbsPath(t *testing.T) {
	// 取得项目绝对路径
	path, err := GetAbsPath()
	fmt.Println(path)
	if err != nil {
		t.Error(err)
	}
	// 检查结果
	target := Projname + "/"
	if path[len(path)-len(target):] != target {
		t.Error("截取错误")
	}
}

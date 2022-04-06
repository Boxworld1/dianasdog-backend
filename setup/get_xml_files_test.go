// @Title  unpackfile
// @Description  Read XML files in folder and unpack them
// @Author  于沛楠
// @Update  2022/3/16
package setup

import (
	"fmt"
	"testing"
)

// test for function: getFiles
func TestGetXMLFiles(t *testing.T) {
	filesPath, err := GetXMLFiles("./testcase")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(filesPath)

	if len(filesPath) != 2 || filesPath[0] != "./testcase/car_test.xml" { // wrong return value
		t.Error("getFile函数读取错误")
	}

	_, err = GetXMLFiles("./cars")
	if err == nil { //wrong error content
		t.Error("read wrong dirpath but not send error")
	}
}

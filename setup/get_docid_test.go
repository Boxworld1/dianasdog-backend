package setup

import (
	"fmt"
	"testing"
)

func TestGetDocid(t *testing.T) {
	abspath, _ := GetAbsPath()
	filesPath, err := GetXMLFiles(abspath + "data/testcase")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(filesPath)

	if len(filesPath) != 2 || filesPath[0] != abspath+"data/testcase/testcase_car.xml" { // wrong return value
		t.Error("getFile函数读取错误")
	}

	_, err = GetXMLFiles("./cars")
	if err == nil { //wrong error content
		t.Error("read wrong dirpath but not send error")
	}

	for _, file := range filesPath {
		// 解码
		fmt.Println("reading " + file)
		itemList, _, _, _, err := UnpackXMLFile(file, "targetResource")
		if err != nil {
			t.Error()
		}
		// 将每一条数据放入数据库中
		for _, item := range itemList {
			docid := GetDocid(item, "targetResource")

			if docid == "" {
				t.Error()
			}
			fmt.Println(docid)
			//StoreItem(item, targetResource, "insert", "docid")
		}
	}
}
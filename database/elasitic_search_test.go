// @title	TestEsInterface
// @description	此函数的用途为检查 es 的接口函数正确性
// @auth	mdy		2022/3
// @param	t		*testing.T	testing 用参数
package database

import (
	"testing"
)

func TestEsInterface(t *testing.T) {
	var err error
	client := EsClient

	// Test insert function of es
	_, err = InsertToEs(client, "100", "布加迪威龙")
	if err != nil {
		t.Error(err)
	}

	// Test update function of es
	_, err = UpdateToEs(client, "100", "柯尼塞格")
	if err != nil {
		t.Error(err)
	}

	// Test search function of es
	_, err = SearchFromEs(client, "柯尼塞格")
	if err != nil {
		t.Error(err)
	}

	// Test fetch all function of es
	_, err = FetchAllFromEs(client)
	if err != nil {
		t.Error(err)
	}

	// Test delete function of es
	DeleteFromES(client, "100")

}

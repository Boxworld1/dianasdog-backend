package database

import (
	"testing"
)

func TestEsInterface(t *testing.T) {
	client, err := ConnectToEs()
	if err != nil {
		t.Error(err)
	} else {
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
}

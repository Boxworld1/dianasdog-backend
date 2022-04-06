package database

import (
	"testing"
)

func TestInsertToEs(t *testing.T) {
	_, err := InsertToEs("100", "布加迪威龙")
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateToEs(t *testing.T) {
	_, err := UpdateToEs("100", "柯尼塞格")
	if err != nil {
		t.Error(err)
	}
}

func TestSearchFromEs(t *testing.T) {
	_, err := SearchFromEs("柯尼塞格")
	if err != nil {
		t.Error(err)
	}
}

func TestFetchAllFromEs(t *testing.T) {
	_, err := FetchAllFromEs()
	if err != nil {
		t.Error(err)
	}
}

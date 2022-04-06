package database

import (
	"testing"
)

func TestInsert(t *testing.T) {
	_, err := Insert("100", "布加迪威龙")
	if err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {
	_, err := Update("100", "柯尼塞格")
	if err != nil {
		t.Error(err)
	}
}

func TestSearch(t *testing.T) {
	_, err := Search("柯尼塞格")
	if err != nil {
		t.Error(err)
	}
}

func TestFetchAll(t *testing.T) {
	_, err := FetchAll()
	if err != nil {
		t.Error(err)
	}
}

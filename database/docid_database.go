// @title	file_database
// @description	本文件函数的用途是配置后端文件对应的数据库
// @auth	ryl		2022/4/20	11:30
// @param	t		*testing.T	testing 用参数

package database

import (
	"database/sql"
	"errors"
)

// 文件数据库接口
var DocidClient *sql.DB

func init() {
	// 创建数据库
	CreateDatabase("docid")

	// 开启数据库
	DocidClient, _ = sql.Open("mysql", GenUrl("docid"))
	inittask := "SET NAMES utf8 "

	// 生成 docid 数据库（每个特型卡只有一个对应文件）
	DocidClient.Exec(inittask)
}

// 新建文件表格（含文件名和内容）
func CreateDocidTable(db *sql.DB, tableName string) error {
	task := "CREATE TABLE IF NOT EXISTS " + tableName + " (docid VARCHAR(64) PRIMARY KEY NULL, data MEDIUMBLOB NULL, filename VARCHAR(64)) DEFAULT CHARSET=utf8;"
	_, err := db.Exec(task)
	return err
}

// 插入文件
func InsertDocid(db *sql.DB, tableName string, docid string, data []byte, filename string) error {
	task := "REPLACE INTO " + tableName + " VALUES(?,?,?)"
	_, err := db.Exec(task, docid, data, filename)
	return err
}

// 取出 docid 对应的值
func GetDocid(db *sql.DB, tableName string, docid string) ([]byte, error) {
	// 按文件名查找
	task := "SELECT docid, data FROM " + tableName + " WHERE docid=?"
	rows, err := db.Query(task, docid)

	// 对应表格不存在
	if err != nil {
		return nil, err
	}

	// 取出数据
	var name string
	var data []byte
	for rows.Next() {
		err = rows.Scan(&name, &data)
		break
	}
	rows.Close()

	// 若数据不符合条件，则返回错误
	if name != docid {
		return nil, errors.New("No data with docid: " + docid)
	}
	return data, err
}

type DataItem struct {
	Name string
	Data []byte
}

// 取出表格下的所有数据
func GetAllDocid(db *sql.DB, tableName string) ([]DataItem, error) {
	// 按文件名查找
	task := "SELECT docid, data FROM " + tableName
	rows, err := db.Query(task)

	// 对应表格不存在
	if err != nil {
		return nil, err
	}

	// 取出数据
	var result []DataItem = make([]DataItem, 0)
	for rows.Next() {
		var docid string
		var data []byte
		err = rows.Scan(&docid, &data)
		result = append(result, DataItem{docid, data})
	}
	rows.Close()

	return result, err
}

// 取出某文件的所有数据
func GetAllDocidByFilename(db *sql.DB, tableName string, filename string) ([]DataItem, error) {
	// 按文件名查找
	task := "SELECT docid, data FROM " + tableName + " WHERE filename=?"
	rows, err := db.Query(task, filename)

	// 对应表格不存在
	if err != nil {
		return nil, err
	}

	// 取出数据
	var result []DataItem = make([]DataItem, 0)
	for rows.Next() {
		var docid string
		var data []byte
		err = rows.Scan(&docid, &data)
		result = append(result, DataItem{docid, data})
	}
	rows.Close()

	return result, err
}

// 删除特定 docid 的文件
func DeleteDocid(db *sql.DB, tableName string, docid string) error {
	// 按文件名查找
	task := "DELETE FROM " + tableName + " WHERE docid=?"
	_, err := db.Exec(task, docid)
	return err
}

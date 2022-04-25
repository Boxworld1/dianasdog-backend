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
var CategoryClient *sql.DB
var DataClient *sql.DB
var ConfigClient *sql.DB
var TemplateClient *sql.DB

func init() {
	// 创建数据库
	CreateDatabase("category")
	CreateDatabase("data")
	CreateDatabase("config")
	CreateDatabase("template")

	// 开启数据库
	CategoryClient, _ = sql.Open("mysql", GenUrl("category"))
	DataClient, _ = sql.Open("mysql", GenUrl("data"))
	ConfigClient, _ = sql.Open("mysql", GenUrl("config"))
	TemplateClient, _ = sql.Open("mysql", GenUrl("template"))

	inittask := "SET NAMES utf8 "

	// 生成特型卡词典
	CategoryClient.Exec(inittask)
	CreateTableFromDict(CategoryClient, "word", []string{"id"})

	// 生成源数据数据库（每个特型卡有多个对应文件）
	DataClient.Exec(inittask)

	// 生成写入行为配置数据库（每个特型卡只有一个对应文件）
	ConfigClient.Exec(inittask)
	CreateFileTable(ConfigClient, "file")

	// 生成模板配置数据库（每个特型卡只有一个对应文件）
	TemplateClient.Exec(inittask)
	CreateFileTable(TemplateClient, "file")
}

// 新建文件表格（含文件名和内容）
func CreateFileTable(db *sql.DB, tableName string) error {
	task := "CREATE TABLE IF NOT EXISTS " + tableName + " (filename VARCHAR(64) PRIMARY KEY NULL, data LONGBLOB NULL) DEFAULT CHARSET=utf8;"
	_, err := db.Exec(task)
	return err
}

// 插入文件
func InsertFile(db *sql.DB, tableName string, filename string, data []byte) error {
	task := "REPLACE INTO " + tableName + " VALUES(?,?)"
	_, err := db.Exec(task, filename, data)
	return err
}

// 取出文件
func GetFile(db *sql.DB, tableName string, filename string) ([]byte, error) {
	// 按文件名查找
	task := "SELECT filename, data FROM " + tableName + " WHERE filename=?"
	rows, err := db.Query(task, filename)

	// 对于文件表不存在
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

	if name != filename {
		return nil, errors.New("No data with filename = " + filename)
	}
	return data, err
}

// @title	dict
// @description	本地词表接口
// @auth	jz		2022/4/25	11:26
// @auth	ryl		2022/4/20	10:30
package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DictClient *sql.DB

//生成URL
func GenUrl(name string) string {
	return SqlUrl + name + "?charset=utf8"
}

//初始化数据库指针
func init() {
	CreateDatabase("dict")
	DictClient, _ = sql.Open("mysql", GenUrl("dict"))
	inittask := `SET NAMES utf8 `
	DictClient.Exec(inittask)
}

//创建数据库
func CreateDatabase(name string) {
	db, _ := sql.Open("mysql", SqlUrl)
	task := "CREATE DATABASE IF NOT EXISTS " + name
	db.Exec(task)
}

//在词典中创建表
func CreateTableInDict(tableName string) error {
	createTask := `CREATE TABLE IF NOT EXISTS ` + tableName + `(docid VARCHAR(100) NULL,field VARCHAR(100) NULL,word VARCHAR(100) NULL)DEFAULT CHARSET=utf8;`
	_, err := DictClient.Exec(createTask)
	return err
}

//从词典中删除表
func DeleteTableFromDict(tableName string) error {
	deleteTask := `DROP TABLE ` + tableName
	_, err := DictClient.Exec(deleteTask)
	return err
}

//获取词典中所有表名(TableName)
func ShowTablesInDict() ([]string, error) {
	task := "select table_name from information_schema.tables where table_schema = 'dict'"
	rows, err := DictClient.Query(task)
	if err != nil {
		return nil, err
	}
	var tableName string
	tables := []string{}
	for rows.Next() {
		err = rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}
	return tables, nil
}

//向表中插入数据
func InsertToDict(tableName string, docid string, field string, word string) error {
	// selectTask := "select 1 from " + tableName + " where docid=? and field=? and word=?"
	// var tmp string
	// err := DictClient.QueryRow(selectTask, docid, field, word).Scan(&tmp)
	// if err == nil && tmp == "1" {
	// 	return nil
	// }
	insertTask := "REPLACE INTO " + tableName + "(docid, field, word) values(?, ?, ?)"
	_, err := DictClient.Exec(insertTask, docid, field, word)
	return err
}

//删除docid为xxx的所有数据
func DeleteByDocidFromDict(tableName string, docid string) error {
	deleteTask := "delete from " + tableName + " where docid=?"
	_, err := DictClient.Exec(deleteTask, docid)
	return err
}

//删除指定docid、field的数据
func DeleteByFieldFromDict(tableName string, docid string, field string) error {
	deleteTask := "delete from " + tableName + " where docid=? and field=?"
	_, err := DictClient.Exec(deleteTask, docid, field)
	return err
}

//返回docid为xxx的所有数据(field + word)
func SearchByDocidFromDict(tableName string, docid string) ([][2]string, error) {
	selectTask := "select field, word from " + tableName + " where docid=?"
	rows, err := DictClient.Query(selectTask, docid)
	if err != nil {
		return nil, err
	}
	var tmp [2]string
	res := [][2]string{}
	for rows.Next() {
		err = rows.Scan(&tmp[0], &tmp[1])
		if err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}
	return res, nil
}

//返回指定docid、field的所有数据(word)
func SearchByFieldFromDict(tableName string, docid string, field string) ([]string, error) {
	selectTask := "select word from " + tableName + " where docid=? and field=?"
	rows, err := DictClient.Query(selectTask, docid, field)
	if err != nil {
		return nil, err
	}
	var tmp string
	res := []string{}
	for rows.Next() {
		err = rows.Scan(&tmp)
		if err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}
	return res, nil
}

//返回表中所有的字段名（去重后的field）
func GetAllFieldFromDict(tableName string) ([]string, error) {
	selectTask := "select distinct field from " + tableName
	rows, err := DictClient.Query(selectTask)
	if err != nil {
		return nil, err
	}
	var word string
	res := []string{}
	for rows.Next() {
		err = rows.Scan(&word)
		if err != nil {
			return nil, err
		}
		res = append(res, word)
	}
	return res, nil
}

//返回field为xxx的所有数据(word)
func GetAllWordFromDict(tableName string, field string) ([]string, error) {
	selectTask := "select word from " + tableName + " where field=?"
	rows, err := DictClient.Query(selectTask, field)
	if err != nil {
		return nil, err
	}
	var word string
	res := []string{}
	for rows.Next() {
		err = rows.Scan(&word)
		if err != nil {
			return nil, err
		}
		res = append(res, word)
	}
	return res, nil
}

package database

import "database/sql"

var PatternClient *sql.DB

func init() {
	CreateDatabase("pattern")
	PatternClient, _ = sql.Open("mysql", GenUrl("pattern"))
	inittask := `SET NAMES utf8 `
	PatternClient.Exec(inittask)
}

//创建表
func CreateTableInPattern(tableName string) error {
	createTask := `CREATE TABLE IF NOT EXISTS ` + tableName + `(pattern VARCHAR(150) NULL primary key)DEFAULT CHARSET=utf8;`
	_, err := PatternClient.Exec(createTask)
	return err
}

//删除表
func DeleteTableFromPattern(tableName string) error {
	deleteTask := `DROP TABLE ` + tableName
	_, err := PatternClient.Exec(deleteTask)
	return err
}

//向表中插入数据
func InsertToPattern(tableName string, pattern string) error {
	insertTask := "Insert ignore into " + tableName + "(pattern)value(?)"
	_, err := PatternClient.Exec(insertTask, pattern)
	return err
}

//从表中删除数据
func DeleteFromPattern(tableName string, pattern string) error {
	deleteTask := "delete from " + tableName + " where pattern=?"
	_, err := PatternClient.Exec(deleteTask, pattern)
	return err
}

//获取表中所有模板
func FetchAllPattern(tableName string) ([]string, error) {
	selectTask := "select * from " + tableName
	rows, err := PatternClient.Query(selectTask)
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

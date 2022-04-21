// @Title: dict
// @Description: reconstruct the local dictionary with mysql and support the insert, search and delete operations
// @Author: 蒋政
// @Update: 2022/3/30
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 词典对应的数据库
var DictClient *sql.DB

func GenUrl(name string) string {
	// var url string = "root:thi4gaiHoa0aicees5booCiet2igoo8i@tcp(mysql-DianasDog.app.secoder.net)/"

	var url string = "root:thi4gaiHoa0aicees5booCiet2igoo8i@tcp(mysql.DianasDog.secoder.local:3306)/"
	// var url string = "root:eelariucie5Tabi8eizioQueineph8la@tcp(localhost:3306)/"
	return url + name + "?charset=utf8"
}

// @title: init
// @description: connect to the default database
// @param: do not need a param
// @return: do not need a return-value
func init() {
	DictClient, _ = sql.Open("mysql", GenUrl("dict"))
	inittask := `SET NAMES utf8 `
	DictClient.Exec(inittask)
}

// @title: CreateTableFromDict
// @description: create the tables needed
// @param: tableName    string    the name of the table to be created
//         columns      []string  the name of the columns to be created in the table
// @return: err         error     nil when the table has been created successfully
func CreateTableFromDict(db *sql.DB, tableName string, columns []string) error {
	createTask := `CREATE TABLE IF NOT EXISTS ` + tableName + `( id VARCHAR(64) PRIMARY KEY NULL`
	for i := 1; i < len(columns); i++ {
		createTask += `,` + columns[i] + ` VARCHAR(64) NULL`
	}
	createTask += `)DEFAULT CHARSET=utf8;`
	_, err := db.Exec(createTask)
	return err
}

// @title: DeleteTableFromDict
// @description: Delete the tables mentioned
// @param: tableName    string    the name of the table to be deleted
// @return: err         error     nil when the table has been deleted successfully
func DeleteTableFromDict(db *sql.DB, tableName string) error {
	deleteTask := `DROP TABLE ` + tableName
	_, err := db.Exec(deleteTask)
	return err
}

// @title: ShowTablesInDict
// @description: get all the tablenames in the dict
// @param: No param is needed.
// @return: tables      []string  store the name of all the tables
//		    err         error     nil when the name of the tables has been read successfully
func ShowTablesInDict(db *sql.DB) ([]string, error) {
	task := "select table_name from information_schema.tables where table_schema = 'dict'"
	rows, err := db.Query(task)
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

// @title: ShowColumnsInTables
// @description: get all the columnnames in the table
// @param: tableName    string    the name of the table to be searched
// @return: columns     []string  store the name of all the columns
//		    err         error     nil when the name of the columns has been read successfully
func ShowColumnsInTable(db *sql.DB, tableName string) ([]string, error) {
	task := "select column_name from information_schema.COLUMNS where table_name='" + tableName + "' and table_schema='dict'"
	rows, err := db.Query(task)
	if err != nil {
		return nil, err
	}
	var columnName string
	columns := []string{}
	for rows.Next() {
		err = rows.Scan(&columnName)
		if err != nil {
			return nil, err
		}
		columns = append(columns, columnName)
	}
	return columns, nil
}

// @title: InsertToDict
// @description: insert word into table
// @param: tableName    string    the name of the target table
//         words        []string  the word to be inserted
// @return: err         error     nil when the word has been inserted into the table successfully
func InsertToDict(db *sql.DB, tableName string, words []string) error {
	insertTask := "REPLACE INTO " + tableName + "(id"
	columns, _ := ShowColumnsInTable(db, tableName)
	for i := 1; i < len(columns); i++ {
		insertTask += "," + columns[i]
	}
	insertTask += ") values(?"
	for i := 0; i < len(columns)-1; i++ {
		insertTask += ",?"
	}
	insertTask += ")"

	var interfaceSlice []interface{} = make([]interface{}, len(words))
	for i, d := range words {
		interfaceSlice[i] = d
	}

	_, err := db.Exec(insertTask, interfaceSlice...)
	if err != nil {
		return err
	}
	return nil
}

// @title: SearchFromDict
// @description: search word from table
// @param: tableName    string    the name of the target table
//         id          string    the id of the data to be searched
// @return: res         []string  the result of the search
//          err         error     nil when the word has been successfully searched in the table
func SearchFromDict(db *sql.DB, tableName string, id string) ([]string, error) {
	//selectTask := "select id from " + tableName + " where id=?"
	columns, _ := ShowColumnsInTable(db, tableName)
	selectTask := "select id"
	for i := 1; i < len(columns); i++ {
		selectTask += ", " + columns[i]
	}
	selectTask += " from " + tableName + " where id=?"
	fmt.Println(selectTask)
	var interfaceSlice []interface{} = make([]interface{}, len(columns))
	var tmp [10]string
	for i := 0; i < len(columns); i++ {
		interfaceSlice[i] = &tmp[i]
	}
	err := db.QueryRow(selectTask, id).Scan(interfaceSlice...)
	if tmp[0] != id {
		return nil, err
	}
	var res []string = make([]string, len(columns))
	for i := 0; i < len(columns); i++ {
		res[i] = tmp[i]
	}
	return res, err
}

// @title: DeleteFromDict
// @description: delete word from table
// @param: tableName  string    the name of the target table
//         id         string    the id of the data to be deleted
// @return:err        error     nil when the word has been deleted from the table successfully
func DeleteFromDict(db *sql.DB, tableName string, id string) error {
	deleteTask := "delete from " + tableName + " where id=?"
	_, err := db.Exec(deleteTask, id)
	if err != nil {
		return err
	}
	return nil
}

// @title: QueryColumn
// @description: Get all the word in column
// @param: table_name    string    the name of the target table
//         column        string    the name of the column
// @return: dictionary    []string    store the total word in the table
//          err           error       nil when the word has been deleted from the table successfully
func QueryColumn(db *sql.DB, tableName string, column string) ([]string, error) {
	QueryTask := "select " + column + " from " + tableName
	rows, err := db.Query(QueryTask)
	if err != nil {
		return nil, err
	}
	var word string
	dictionary := []string{}
	for rows.Next() {
		err = rows.Scan(&word)
		if err != nil {
			return nil, err
		}
		dictionary = append(dictionary, word)
	}
	return dictionary, nil
}

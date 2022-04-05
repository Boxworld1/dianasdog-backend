// @Title:  dict
// @Description: reconstruct the local dictionary with mysql and support the insert, search and delete operations
// @Author:  蒋政
// @Update:  2022/3/30
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var dataSourceName string = "root:thi4gaiHoa0aicees5booCiet2igoo8i@tcp(mysql.DianasDog.secoder.local:3306)/dict?charset=utf8"

// @title:	init
// @description: connect to the default database
// @param: do not need in-params
// @return: do not need a return-value
func init() {
	database, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
	}
	db = database
}

// @title:	CreateTableFromDict
// @description: create the tables needed
// @param:	tableName	string		the name of the table to be created
// @return: err			error		nil when the table has been created successfully
func CreateTableFromDict(tableName string) error {
	createTask := `CREATE TABLE IF NOT EXISTS ` + tableName + `(
		word VARCHAR(64) PRIMARY KEY NULL
	)DEFAULT CHARSET=utf8;
	`
	_, err := db.Exec(createTask)
	return err
}

// @title:	DeleteTableFromDict
// @description: Delete the tables mentioned
// @param:	tableName	string		the name of the table to be deleted
// @return: err			error			nil when the table has been deleted successfully
func DeleteTableFromDict(tableName string) error {
	deleteTask := `DROP TABLE ` + tableName
	_, err := db.Exec(deleteTask)
	return err
}

// @title:	InsertToDict
// @description:  insert word into table
// @param:	tableName   string      the name of the target table
//			word		string		the word to be inserted
// @return: err         error      nil when the word has been inserted into the table successfully
func InsertToDict(tableName string, word string) error {
	insertTask := "INSERT IGNORE INTO " + tableName + "(word) values(?)"
	stmt, err := db.Prepare(insertTask)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(word)
	if err != nil {
		return err
	}
	return nil
}

// @title:   SearchFromDict
// @description:  search word from table
// @param:	tableName   string      the name of the target table
//			word		string		the word to be searched
// @return: word		 string		the name of the word if it is in the table else "None"
//			err          error      nil when the word is in the table
func SearchFromDict(tableName string, word string) (string, error) {
	selectTask := "select word from " + tableName + " where word=?"
	var res string
	err := db.QueryRow(selectTask, word).Scan(&res)
	if err == nil && res == word {
		return res, err
	} else {
		return "None", err
	}
}

// @title:	DeleteFromDict
// @description:	delete word from table
// @param:	word		 string		the word to be deleted
//			table_name   string     the name of the target table
// @return: err          error      nil when the word has been deleted from the table successfully
func DeleteFromDict(tableName string, word string) error {
	deleteTask := "delete from " + tableName + " where word=?"
	stmt, err := db.Prepare(deleteTask)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(word)
	if err != nil {
		return err
	}
	return nil
}

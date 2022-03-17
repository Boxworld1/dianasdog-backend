// @Title  dict
// @Description  construct the local dictionary and support the insert, search and delete operations
// @Author  蒋政
// @Update  2022/3/16
package dict

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var dataSourceName string = "./dict.db"

// @title   CreateTable
// @description  create the tables needed
// @auth    蒋政       2022/3/16
// @param	tableName   string             the name of the table to be created
// @return  err          error              nil when the table has been created successfully
func CreateTable(tableName string) error {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return err
	}
	//set the table
	sqlTable := `CREATE TABLE IF NOT EXISTS ` + tableName + `(
		word VARCHAR(64) PRIMARY KEY NULL
	);
	`
	db.Exec(sqlTable)
	db.Close()
	return nil
}

// @title   Insert
// @description  insert word into table
// @auth    蒋政       2022/3/16
// @param	word		 string				the word to be inserted
//			tableName   string             the name of the target table
// @return  err          error              nil when the word has been inserted into the table successfully
func Insert(word string, tableName string) error {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return err
	}
	//judge whether the word has been saved
	tmp, _ := Search(word, tableName)
	if tmp == word {
		return nil
	} else {

		//insert word into table
		insertTask := "INSERT INTO " + tableName + "(word) values(?)"
		stmt, err := db.Prepare(insertTask)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(word)
		if err != nil {
			return err
		}
		db.Close()
		return nil
	}
}

// @title   Search
// @description  search word in table
// @auth    蒋政       2022/3/16
// @param	wordName		 string				the word to be searched
//			tableName   string             the name of the target table
// @return  word		 string				the name of the word if it is in the table else "None"
//			err          error              nil when the word is in the table
func Search(wordName string, tableName string) (string, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return "None", err
	}
	//search word in table
	searchTask := "SELECT * FROM " + tableName
	rows, err := db.Query(searchTask)
	if err != nil {
		return "None", err
	}
	var word string
	for rows.Next() {
		err = rows.Scan(&word)
		if err != nil {
			return "None", err
		}
		if word == wordName {
			rows.Close()
			return wordName, nil
		}
	}
	rows.Close()
	db.Close()
	return "None", err
}

// @title   Delete
// @description  delete word from table
// @auth    蒋政       2022/3/16
// @param	word		 string				the word to be deleted
//			table_name   string             the name of the target table
// @return  err          error              nil when the word has been deleted from the table successfully
func Delete(word string, tableName string) error {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return err
	}
	//delete word from table
	deleteTask := "delete from " + tableName + " where word=?"
	stmt, err := db.Prepare(deleteTask)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(word)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

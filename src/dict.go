// @Title  dict
// @Description  construct the local dictionary and support the insert, search and delete operations
// @Author  蒋政
// @Update  2022/3/16
package dict

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// @title   CreateTable
// @description  create the tables needed
// @auth    蒋政       2022/3/16
// @param	table_name   string             the name of the table to be created
// @return  err          error              nil when the table has been created successfully
func CreateTable(table_name string) error {
	db, err := sql.Open("sqlite3", "./data/dict.db")
	if err != nil {
		return err
	}
	//set the table
	sql_table := `CREATE TABLE IF NOT EXISTS ` + table_name + `(
		word VARCHAR(64) PRIMARY KEY NULL
	);
	`
	db.Exec(sql_table)
	db.Close()
	return nil
}

// @title   Insert
// @description  insert word into table
// @auth    蒋政       2022/3/16
// @param	word		 string				the word to be inserted
//			table_name   string             the name of the target table
// @return  err          error              nil when the word has been inserted into the table successfully
func Insert(word string, table_name string) error {
	db, err := sql.Open("sqlite3", "./dict.db")
	if err != nil {
		return err
	}
	//judge whether the word has been saved
	tmp, _ := Search(word, table_name)
	if tmp == word {
		return nil
	} else {

		//insert word into table
		insert_task := "INSERT INTO " + table_name + "(word) values(?)"
		stmt, err := db.Prepare(insert_task)
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
// @param	word_name		 string				the word to be searched
//			table_name   string             the name of the target table
// @return  word		 string				the name of the word if it is in the table else "None"
//			err          error              nil when the word is in the table
func Search(word_name string, table_name string) (string, error) {
	db, err := sql.Open("sqlite3", "dict.db")
	if err != nil {
		return "None", err
	}
	//search word in table
	search_task := "SELECT * FROM " + table_name
	rows, err := db.Query(search_task)
	if err != nil {
		return "None", err
	}
	var word string
	for rows.Next() {
		err = rows.Scan(&word)
		if err != nil {
			return "None", err
		}
		if word == word_name {
			rows.Close()
			return word_name, nil
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
func Delete(word string, table_name string) error {
	db, err := sql.Open("sqlite3", "dict.db")
	if err != nil {
		return err
	}
	//delete word from table
	delete_task := "delete from " + table_name + " where word=?"
	stmt, err := db.Prepare(delete_task)
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

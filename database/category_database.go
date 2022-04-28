// @title	file_database
// @description	本文件函数的用途是配置类別对应的数据库
// @auth	ryl		2022/4/20	11:30
// @param	t		*testing.T	testing 用参数

package database

import (
	"database/sql"
)

// 文件数据库接口
var CategoryClient *sql.DB

func init() {
	// 创建数据库
	CreateDatabase("category")

	// 开启数据库
	CategoryClient, _ = sql.Open("mysql", GenUrl("category"))

	// 生成特型卡数据库
	CategoryClient.Exec("SET NAMES utf8 ")
	CreateCategoryTable(CategoryClient, "word")
}

// 新建类名表格
func CreateCategoryTable(db *sql.DB, tableName string) error {
	task := "CREATE TABLE IF NOT EXISTS " + tableName + " (category VARCHAR(64) PRIMARY KEY NULL) DEFAULT CHARSET=utf8;"
	_, err := db.Exec(task)
	return err
}

// 插入类名
func InsertCategory(db *sql.DB, tableName string, category string) error {
	task := "REPLACE INTO " + tableName + " VALUES(?)"
	_, err := db.Exec(task, category)
	return err
}

// 取出所有类名
func GetAllCategory(db *sql.DB, tableName string) ([]string, error) {
	// 查找表格
	task := "SELECT category FROM " + tableName
	rows, err := db.Query(task)

	// 对应表格不存在
	if err != nil {
		return nil, err
	}

	// 否则取出数据
	var names []string = make([]string, 0)
	for rows.Next() {
		var name string
		rows.Scan(&name)
		// 若不为测试类型则加入
		if name != "testdata" && name != "testcase_car" {
			names = append(names, name)
		}
	}
	rows.Close()

	return names, nil
}

// 统计个数
func CountCategory(db *sql.DB, tableName string) (int, error) {
	// 查找表格
	task := "SELECT count(*) FROM " + tableName
	rows, err := db.Query(task)

	// 对应表格不存在
	if err != nil {
		return 0, err
	}

	// 否则取出数据
	var count int = 0
	for rows.Next() {
		rows.Scan(&count)
		break
	}
	rows.Close()

	return count, nil
}

// 删除表格
func DropCategory(db *sql.DB, tableName string) error {
	// 查找表格
	task := "DROP TABLE " + tableName
	_, err := db.Query(task)

	// 对应表格不存在
	return err
}

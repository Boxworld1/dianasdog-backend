package database

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Level    string `json:"level"`
}

// EncodePassword
// @title:	EncodePassword
// @description: 加密一个密码
// @param: password string 要加密的密码
// @return: encodePWD,err string,error 加密后的密码,错误信息
func EncodePassword(password string) (string, error) {
	fmt.Println("正在加密")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return "", err
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	fmt.Println("加密完成")
	return encodePWD, err
}

// @title:	init
// @description: connect to the default database
// @param: do not need in-params
// @return: do not need a return-value
//func Init() error {
//	database, err := sql.Open("mysql", dataSourceName)
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//	UserInfoClient = database
//	return nil
//}

var UserInfoClient *sql.DB

func init() {
	CreateDatabase("UserInfo")
	UserInfoClient, _ = sql.Open("mysql", GenUrl("UserInfo"))
	inittask := `SET NAMES utf8 `
	UserInfoClient.Exec(inittask)
	_ = CreateTableForUserinfo()
	_ = CreateTableForUserLevel()
}

// CreateTableForUserinfo
// @title:	CreateTableForUserinfo
// @description: 建立存储user用户名和密码的表
// @param: do not need in-params
// @return: err 错误信息
func CreateTableForUserinfo() error {
	createTask := `CREATE TABLE IF NOT EXISTS ` + "UserInfo" + `( 
		username VARCHAR(100) not null, userpassword VARCHAR(500) not null,  PRIMARY KEY (username)
	)DEFAULT CHARSET=utf8;
	`
	_, err := UserInfoClient.Exec(createTask)
	return err
}

// CreateTableForUserLevel
// @title:	CreateTableForUserLevel
// @description: 建立存储user用户名和权限等级的表
// @param: do not need in-params
// @return: err 错误信息
func CreateTableForUserLevel() error {
	createTask := `CREATE TABLE IF NOT EXISTS ` + "UserLevel" + `( 
		username VARCHAR(100) not null, userlevel VARCHAR(10) not null,  PRIMARY KEY (username)
	)DEFAULT CHARSET=utf8;
	`
	_, err := UserInfoClient.Exec(createTask)
	return err
}

// InsertPwdIntoSQL
// @title:	InsertPwdIntoSQL
// @description: 向数据库中添加一个用户的信息
// @param: encodedPassword,username,userlevel string,string,string 分别是用户的密码，用户名，权限等级
// @return: err 错误信息
func InsertPwdIntoSQL(encodedPassword string, username string, userlevel string) error {
	fmt.Println("正在将一条用户信息插入sql")
	insertTask := "INSERT IGNORE INTO " + "UserInfo" + "(username, userpassword) values('" + username + "','" + encodedPassword + "')"
	_, err := UserInfoClient.Exec(insertTask)
	if err != nil {
		return err
	}
	insertTask = "INSERT IGNORE INTO " + "UserLevel" + "(username, userlevel) values('" + username + "','" + userlevel + "')"
	_, err = UserInfoClient.Exec(insertTask)
	if err != nil {
		return err
	}
	fmt.Println("插入完毕")
	return nil
}

// UserSignup
// @title:	UserSignup
// @description: 向数据库中添加一个用户的信息
// @param: user User 一个封装好的用户信息
// @return: err 错误信息
func UserSignup(user User) error {
	fmt.Println("注册用户信息")
	//err := Init()
	//if err != nil {
	//	return err
	//}
	_ = CreateTableForUserinfo()
	err := InsertPwdIntoSQL(user.Password, user.Name, user.Level)
	if err != nil {
		return err
	}
	return nil //和前端商量一下，可能要返回个码
}

// SearchUser
// @title:	SearchUser
// @description: 根据用户名查找一个用户的密码和权限等级
// @param: username string  用户名
// @return: password, level, err  string, string, error 分别是密码，权限等级和错误信息
func SearchUser(username string) (string, string, error) {
	selectTask := "select userpassword from UserInfo" + " where username='" + username + "'"
	res := UserInfoClient.QueryRow(selectTask)
	var password string
	var level string
	err := res.Scan(&password)
	selectTask = "select userlevel from UserLevel" + " where username='" + username + "'"
	res = UserInfoClient.QueryRow(selectTask)
	err = res.Scan(&level)

	if err == nil {
		return password, level, err
	} else {
		return "None", "None", err
	}

}

// UserSignIn
// @title:	UserSignIn
// @description: 根据用户名查找一个用户的密码和权限等级
// @param: username string  用户名
// @return: EncodedPassword, level, err  string, string, error 分别是密码，权限等级和错误信息
func UserSignIn(username string) (string, string, error) {
	fmt.Println("验证用户信息")
	//err := Init()
	//if err != nil {
	//	return "None", err
	//}
	EncodedPassword, level, err1 := SearchUser(username)
	if err1 != nil {
		return "None", "None", err1
	}
	return EncodedPassword, level, nil
}

// DeleteUser
// @title:	DeleteUser
// @description: 根据用户名删除一个用户
// @param: username string  用户名
// @return: err  error 错误信息
func DeleteUser(username string) error {
	fmt.Println("正在删除一条用户信息")
	deleteTask := "DELETE FROM " + "UserInfo" + " where username='" + username + "'"
	_, err := UserInfoClient.Exec(deleteTask)
	if err != nil {
		return err
	}
	deleteTask = "DELETE FROM " + "UserLevel" + " where username='" + username + "'"
	_, err = UserInfoClient.Exec(deleteTask)
	if err != nil {
		return err
	}
	fmt.Println("删除完毕")
	return nil
}

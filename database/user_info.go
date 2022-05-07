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
}

func CreateTableForUserinfo() error {
	createTask := `CREATE TABLE IF NOT EXISTS ` + "UserInfo" + `( 
		username VARCHAR(100) not null, userpassword VARCHAR(500) not null, userlevel VARCHAR(10) not null,  PRIMARY KEY (username)
	)DEFAULT CHARSET=utf8;
	`
	_, err := UserInfoClient.Exec(createTask)
	return err
}

func InsertPwdIntoSQL(encodedPassword string, username string, userlevel string) error {
	fmt.Println("正在将一条用户信息插入sql")
	insertTask := "INSERT IGNORE INTO " + "UserInfo" + "(username, userpassword, userlevel) values('" + username + "','" + encodedPassword + "','" + userlevel + "')"
	_, err := UserInfoClient.Exec(insertTask)
	if err != nil {
		return err
	}
	fmt.Println("插入完毕")
	return nil
}

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

func SearchUser(username string) (string, error) {
	selectTask := "select userpassword from UserInfo" + " where username='" + username + "'"
	res := UserInfoClient.QueryRow(selectTask)
	var password string
	err := res.Scan(&password)
	if err == nil {
		return password, err
	} else {
		return "None", err
	}

}

func UserSignIn(username string) (string, error) {
	fmt.Println("验证用户信息")
	//err := Init()
	//if err != nil {
	//	return "None", err
	//}
	EncodedPassword, err1 := SearchUser(username)
	if err1 != nil {
		return "None", err1
	}
	return EncodedPassword, nil
}

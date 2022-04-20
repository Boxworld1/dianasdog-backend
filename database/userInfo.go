package database

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
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
	UserInfoClient, _ = sql.Open("mysql", GenUrl("UserInfo"))
	inittask := `SET NAMES utf8 `
	UserInfoClient.Exec(inittask)
	_ = CreateTableForUserinfo()
}

func CreateTableForUserinfo() error {
	createTask := `CREATE TABLE IF NOT EXISTS ` + "UserInfo" + `( 
		username VARCHAR(100) not null, userpassword VARCHAR(500) not null,  PRIMARY KEY (username)
	)DEFAULT CHARSET=utf8;
	`
	_, err := UserInfoClient.Exec(createTask)
	return err
}

func InsertPwdIntoSQL(encodedPassword string, username string) error {
	fmt.Println("正在将一条用户信息插入sql")
	insertTask := "INSERT IGNORE INTO " + "UserInfo" + "(username, userpassword) values('" + username + "','" + encodedPassword + "')"
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
	err := InsertPwdIntoSQL(user.Password, user.Name)
	if err != nil {
		return err
	}
	return nil //和前端商量一下，可能要返回个码
}

func SearchUser(username string) (string, error) {

	selectTask := "select username,userpassword from UserInfo" + " where username='" + username + "'"
	var res string
	var tmp string
	err := UserInfoClient.QueryRow(selectTask, username).Scan(&tmp, &res)
	if err == nil {
		return res, err
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

package database

import (
	"testing"
)

// test for function: EncodePassword
func TestEncodePassword(t *testing.T) {
	_, err := EncodePassword("pwd")
	if err != nil {
		t.Error(err)
	}
}

// test for function: CreateTableForUserinfo
func TestCreateTableForUserinfo(t *testing.T) {
	err := CreateTableForUserinfo()
	if err != nil {
		t.Error(err)
	}
}

// test for function: UserSignup
func TestUserSignup(t *testing.T) {
	var user User
	user.Name = "tester"
	user.Password = "pwd"
	user.Level = "3"
	err := UserSignup(user)
	if err != nil {
		t.Error(err)
	}
}

// test for function: SearchUser
func TestSearchUser(t *testing.T) {
	password, level, err := SearchUser("tester")
	if password != "pwd" || level != "3" {
		t.Error(password)
		t.Error("查询的密码错误")
	}
	if err != nil {
		t.Error(err)
	}
}

// test for function: UserSignIn
func TestUserSignIn(t *testing.T) {
	res, level, err := UserSignIn("tester")
	if res != "pwd" || level != "3" {
		t.Error("登陆失败")
	}
	if err != nil {
		t.Error(err)
	}
}

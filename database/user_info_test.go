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
	user.Name = "tester1"
	user.Password = "pwd"
	user.Level = "3"
	err := UserSignup(user)
	if err != nil {
		t.Error(err)
	}
}

// test for function: SearchUser
func TestSearchUser(t *testing.T) {
	password, level, err := SearchUser("tester1")
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
	res, level, err := UserSignIn("tester1")
	if res != "pwd" || level != "3" {
		t.Error("登陆失败")
	}
	if err != nil {
		t.Error(err)
	}
}

// test for function: DeleteUser
func TestDeleteUser(t *testing.T) {
	err := DeleteUser("tester1")
	if err != nil {
		t.Error(err)
	}
}

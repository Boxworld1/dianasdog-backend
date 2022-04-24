// @title		GetAbsPath
// @description	取得项目的绝对路径
// @auth		ryl				2022/4/7		12:00
// @return		path			string			绝对路径
// @return		err				error			错误值

package path

import (
	"os"
	"strings"
)

var Projname = "Backend-Test"

func GetAbsPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path = strings.Replace(path, Projname, Projname+"@", 1)
	pathlist := strings.Split(path, "@")

	return pathlist[0], nil
}

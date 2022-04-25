// @title	SetTestData
// @description	将单元测试需要的数据入库
// @auth	ryl			2022/4/25		18:00
// @return	err			error			错误值

package io

import (
	"dianasdog/database"
	"dianasdog/path"
	"io/ioutil"
)

func SetTestData() error {

	// 得到此文件的绝对路径
	abspath, _ := path.GetAbsPath()

	// 数据加入数据库
	res := "testdata"
	filename := "testcase.xml"
	filepath := abspath + "testcase/" + filename

	// 读入文件
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	// 若无对应表则建表
	if err := database.CreateFileTable(database.DataClient, res); err != nil {
		return err
	}
	if err := database.InsertFile(database.DataClient, res, filename, data); err != nil {
		return err
	}

	// 数据加入数据库
	filename = "config.json"
	filepath = abspath + "testcase/" + filename

	// 读入文件
	data, err = ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	// 加入数据库中
	if err := database.InsertFile(database.ConfigClient, "file", res, data); err != nil {
		return err
	}

	// 数据加入数据库
	filename = "template.json"
	filepath = abspath + "testcase/" + filename

	// 读入文件
	data, err = ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	// 加入数据库中
	if err := database.InsertFile(database.TemplateClient, "file", res, data); err != nil {
		return err
	}

	return nil
}

// @title		SetByCategory
// @description	按配置文件中的特型卡表分別读入数据
// @auth		ryl				2022/4/7		10:00
// @return		err				error			错误值

package setup

import (
	"bufio"
	"dianasdog/path"
	"fmt"
	"os"
)

func SetByCategory() error {
	var category []string

	// 得到此文件的绝对路径
	abspath, _ := path.GetAbsPath()

	// 以文件形式读入特型卡类型
	catfile, err := os.Open(abspath + "config/category.txt")

	// 出现错误则退出
	if err != nil {
		fmt.Println(err)
		return err
	}

	fileScanner := bufio.NewScanner(catfile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		category = append(category, fileScanner.Text())
	}

	catfile.Close()
	// 遍历特型卡名字文档
	for _, cat := range category {
		// 设置相对路径
		path := abspath + "data/"
		// 调用增加数据接口
		err := AddData(path+cat+"/", cat)
		if err != nil {
			return err
		}
	}

	return nil
}

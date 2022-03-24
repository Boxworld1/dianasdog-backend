// @title		IntentRecognition
// @description	此函数的用途为意图识別，需要找到匹配字符串的核心词并向下传递
// @auth		ryl				2022/3/24		11:00
// @param		targetString	string			目标句

package search_service

import "fmt"

func IntentRecognition(targetString string) []string {

	var resources = make([]string, 0)
	len := len(targetString)

	fmt.Println(len)

	// 利用词表接口搜索，每次存入不同长度字符串的后缀查找
	for cnt := 0; cnt < len; cnt++ {
		fmt.Println(cnt)
		fmt.Println(targetString[cnt:len])
		result := WordSearch(targetString[cnt:len])
		resources = append(resources, result...)
	}

	return resources
}

// @title	Search
// @description	此函数的用途为搜索与句子相关的信息
// @auth	ryl		2022/4/26	22:00
// @param	query	string		句子
// @return	err		string		结果

package search

import "fmt"

func Search(query string) string {

	// 意图识別
	intentList := IntentionRecognition(query)

	// Query 理解
	result := QueryUnderstanding(intentList, query)

	for item := range result {
		fmt.Println(item)
	}

	// 返回搜索结果
	return "I'm result"
}

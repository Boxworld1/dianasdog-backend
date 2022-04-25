// @title	Search
// @description	此函数的用途为搜索与句子相关的信息
// @auth	ryl		2022/4/25	22:00
// @param	query	string		句子
// @return	err		string		结果

package search

func Search(query string) string {

	// 意图识別
	intentList := IntentionRecognition(query)

	result := QueryUnderstanding(intentList, query)

	for item := range result {
		item = item
	}

	// 无论正确与否都返回 err 的内容
	return ""
}

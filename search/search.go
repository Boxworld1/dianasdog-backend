// @title	Search
// @description	此函数的用途为搜索与句子相关的信息
// @auth	ryl		2022/4/27	16:00
// @param	query	string		句子
// @return	resList	[]string	结果

package search

import "dianasdog/database"

func Search(query string) []string {

	// 意图识別
	intentList := IntentionRecognition(query)

	// Query 理解
	result := QueryUnderstanding(intentList, query)

	// es搜索
	var docIdList []string

	for item := range result {
		resourceName := result[item].Resource
		var content string
		content = ""
		for x := range result[item].detail {
			if result[item].pattern[x] != "garbage" && result[item].pattern[x] != "intent" {
				content += result[item].detail[x]
				if x < len(result[item].detail) {
					content += " "
				}
			}
		}
		resList, _ := database.SearchFromEs(resourceName, database.EsClient, content)
		for k := range resList {
			docIdList = append(docIdList, resList[k].DocID)
		}
	}

	var resList []string = make([]string, 0)
	for _, docid := range docIdList {
		result, _ := database.GetFromRedis(database.RedisClient, docid)
		resList = append(resList, result)
	}

	// 返回搜索结果
	return resList
}

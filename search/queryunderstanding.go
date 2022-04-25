package search

import (
	"dianasdog/database"
	"strings"
)

var f [55][15]int
var w [55][55][15]int
var ff [55][15][15]string

type result struct {
	Resource string
	pattern  []string
	detail   []string
}

func QueryUnderstanding(intentlist []string, query string) []result {
	var retList []result
	for _, table := range intentlist {
		patterns, _ := database.FetchAllPattern(table)
		for _, rawpattern := range patterns {
			pattern := strings.Split(rawpattern, "+")
			lenp, lenq := len(pattern), len([]rune(query))
			rune := []rune(query)
			//初始化数组
			for i := 0; i <= lenq; i++ {
				for j := 0; j <= lenp; j++ {
					f[i][j] = 0
					for k := 0; k <= lenp; k++ {
						ff[i][j][k] = ""
					}
				}
				for j := 0; j <= lenq; j++ {
					for k := 0; k <= lenp; k++ {
						w[i][j][k] = 0
					}
				}
			}
			f[0][0] = 1
			//检测query[i:j]是否在模板对应的词表中
			for i, field := range pattern {
				rets := acmap[table][field].Match(query)
				for _, ret := range rets {
					w[ret.BegPosition+1][ret.EndPosition+1][i+1] = 1
				}
			}
			//动态规划
			for i := 1; i <= lenq; i++ {
				for j := 1; j <= lenp; j++ {
					for k := i - 1; k >= 0; k-- {
						if f[k][j-1] == 1 && w[k+1][i][j] == 1 {
							f[i][j] = 1
							ff[i][j] = ff[k][j-1]
							ff[i][j][j] = string(rune[k:i])
							break
						}
					}
				}
			}
			//返回结果
			if f[lenq][lenp] == 1 {
				det := []string{}
				for i := 1; i <= lenp; i++ {
					det = append(det, ff[lenq][lenp][i])
				}
				retList = append(retList, result{table, pattern, det})
			}
		}
	}
	return retList
}

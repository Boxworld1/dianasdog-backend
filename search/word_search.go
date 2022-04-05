// @title		WordSearch
// @description	利用 Trie 树回传匹配成功的特形卡类型
// @auth		ryl				2022/3/24		11:00
// @param		targetResource	string			特型卡片类型

package search

func WordSearch(targetString string) []string {
	var result = make([]string, 0)

	len := len(targetString)
	// 模拟 Trie 树，后续应与数据库交互！
	if len < 1 {
		return result
	}

	if targetString[0] == 'c' {

		if len < 2 {
			return result
		}

		if targetString[1] == 'a' {

			if len < 3 {
				return result
			}
			if targetString[2] == 'r' {
				result = append(result, "car")
			}
			if targetString[2] == 'd' {
				result = append(result, "cad")
			}

		}

		if targetString[1] == 'd' {
			result = append(result, "cd")
		}
	}

	return result
}

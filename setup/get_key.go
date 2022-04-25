// @title	GetKey
// @description	此函数的用途为得到数据的 key
// @auth	wzq		2022/4/25		21:05
// @param	item	*etree.Element	数据 item
// @param	targetResource	string	特型卡 Id
// @return	docid	string			docid

package setup

import "strings"

func GetKey(tarPath string) string {
	path := strings.Split(tarPath, ".")
	key := path[len(path)-1]
	return key
}

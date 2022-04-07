// @title   GetXMLFiles
// @description  get XML files from a certain folder
// @auth    于沛楠		2022/3/16
// @param	dirPath		string  	name of folder which contains XML files
// @return  filePath	[]string	path array of XML files under dirPath
// @return	err			error 		non-nil when dirPath is wrong

package setup

import (
	"io/ioutil"

	"strings"
)

func GetXMLFiles(dirPath string) (filesPath []string, err error) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	for _, file := range dir {
		ok := strings.HasSuffix(file.Name(), ".xml")
		if ok {
			filesPath = append(filesPath, dirPath+"/"+file.Name())
		}
	}
	return filesPath, err
}

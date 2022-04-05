// @Title  unpackfile
// @Description  Read XML files in folder and unpack them
// @Author  于沛楠
// @Update  2022/3/16
package setup

import (
	"io/ioutil"

	"strings"
)

// @title   GetXMLFiles
// @description  get XML files from a certain folder
// @auth    于沛楠     2022/3/16
// @param	dirPath    string       name of folder which contains XML files
// @return  filePath   []string   	path array of XML files under dirPath
//		    err        error        non-nil when dirPath is wrong
func GetXMLFiles(dirPath string) (filesPath []string, err error) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	//PthSep := string(os.PathSeparator)
	for _, file := range dir {
		ok := strings.HasSuffix(file.Name(), ".xml")
		if ok {
			filesPath = append(filesPath, dirPath+"/"+file.Name())
		}
	}
	return filesPath, err
}

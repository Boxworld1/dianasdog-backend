package unpack

import (

	"io/ioutil"

	"strings"

	"github.com/beevik/etree"
)

var resources []string

func GetFiles(dirPath string) (filesPath []string, err error) {
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

func Unpack(fileName string, resourceName string) (err error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(fileName); err != nil {
		return err
	}
	root := doc.SelectElement("DOCUMENT")
	cnt := 0
	for i, _ := range root.SelectElements("item") {
		cnt++
		i++
		//调用写入层接口
		// Store(item, resourceName, "insert")

	}

	return nil
}
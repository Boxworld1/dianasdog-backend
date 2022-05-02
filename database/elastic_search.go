// @title	elastic_search
// @description	倒排引擎接口
// @auth	wzq		2022/3
// @auth	ryl		2022/4/20	10:30
// @auth	wzq		2022/4/26   20:34

package database

import (
	"context"
	"fmt"
	"reflect"

	"github.com/olivere/elastic/v7"
)

// 预加载数据库
var EsClient *elastic.Client

// 格式
type Doc struct {
	DocID   string `json:"DocID"`
	Content string `json:"content"`
}

// 全局初始化
func init() {
	EsClient, _ = elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(EsUrl),
	)
}

// InsertToEs
// @title		InsertToEs
// @description	向指定的特型（resourceName）中添加数据
// @auth		wzq				2022/4/26		20:39
func InsertToEs(resourceName string, client *elastic.Client, docId string, content string) (string, error) {
	doc := Doc{DocID: docId, Content: content}
	put1, err := client.Index().
		Index(resourceName).
		BodyJson(doc).
		Id(doc.DocID).
		Do(context.Background())
	if err != nil {
		print(err.Error())
	}
	fmt.Printf("Indexed DocId %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return put1.Id, err
}

// UpdateToEs
// @title		UpdateToEs
// @description	向指定的特型（resourceName）中指定的文档更新数据
// @auth		wzq				2022/4/26		20:39
func UpdateToEs(resourceName string, client *elastic.Client, docId string, newContent string) (string, error) {
	put2, err := client.Update().
		Index(resourceName).
		Id(docId).
		Doc(map[string]interface{}{"content": newContent}).
		Do(context.Background())
	if err != nil {
		print(err.Error())
	}
	fmt.Printf("update content %s\n", put2.Result)
	return put2.Result, err
}

//按照内容去查找，不是精确查找，只要有匹配词就可以
func SearchFromEs(resourceName string, client *elastic.Client, content string) ([]Doc, error) {
	var typ Doc
	var err error
	var put4 *elastic.SearchResult
	matchQuery := elastic.NewMatchPhraseQuery("content", content)
	put4, err = client.Search(resourceName).Query(matchQuery).Size(70).Do(context.Background())
	if err != nil {
		print(err.Error())
		return nil, err
	}
	var result []Doc
	for _, item := range put4.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法，一条条把数据列出来
		t := item.(Doc)
		fmt.Printf("%#v\n", t)
		result = append(result, t)
	}
	return result, err
}

func FetchAllFromEs(resourceName string, client *elastic.Client) ([]Doc, error) { //拿到类型里的所有数据
	var put3 *elastic.SearchResult
	var err error
	//取所有
	put3, err = client.Search(resourceName).Do(context.Background())
	var typ Doc
	if err != nil {
		print(err.Error())
		return nil, err
	}
	var result []Doc
	for _, item := range put3.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Doc)
		fmt.Printf("%#v\n", t)
		result = append(result, t)
	}
	return result, err
}

func DeleteFromES(resourceName string, client *elastic.Client, docId string) { //指定想要删除的文档的docId
	var err error
	res, err := client.Delete().Index(resourceName).
		Id(docId).
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

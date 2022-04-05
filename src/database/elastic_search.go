package database

import (
	"context"
	"fmt"
	"reflect"

	"github.com/olivere/elastic/v7"
)

// Elasticsearch demo

type Doc struct {
	DocID   string `json:"DocID"`
	Content string `json:"content"`
}

func InsertToEs(docId string, content string) (string, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://elasticsearch.DianasDog.secoder.local:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("connect to es success")
	doc := Doc{DocID: docId, Content: content}
	put1, err := client.Index().
		Index("document").
		BodyJson(doc).
		Id(doc.DocID).
		Do(context.Background())
	if err != nil {
		print(err.Error())
	}
	fmt.Printf("Indexed DocId %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return put1.Id, err
}

func UpdateToEs(docId string, newContent string) (string, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://elasticsearch.DianasDog.secoder.local:9200"))
	//下面是更新项目的函数，需要传入docid，对数据进行改变
	put2, err := client.Update().
		Index("document").
		Id(docId).
		Doc(map[string]interface{}{"content": newContent}).
		Do(context.Background())
	if err != nil {
		print(err.Error())

	}
	fmt.Printf("update content %s\n", put2.Result)
	return put2.Result, err
}

func SearchFromEs(content string) ([]Doc, error) { //按照内容去查找，不是精确查找，只要有匹配词就可以
	client, _ := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://elasticsearch.DianasDog.secoder.local:9200"))
	var typ Doc
	var err error
	var put4 *elastic.SearchResult
	matchPhraseQuery := elastic.NewMatchPhraseQuery("content", content)
	put4, err = client.Search("document").Query(matchPhraseQuery).Do(context.Background())
	if err != nil {
		print(err.Error())
	}
	var result []Doc
	for _, item := range put4.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法，一条条把数据列出来
		t := item.(Doc)
		fmt.Printf("%#v\n", t)
		result = append(result, t)
	}
	return result, err
}

func FetchAllFromEs() ([]Doc, error) { //拿到类型document里的所有数据
	client, _ := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://elasticsearch.DianasDog.secoder.local:9200"))
	var put3 *elastic.SearchResult
	var err error
	//取所有
	put3, err = client.Search("document").Do(context.Background())
	var typ Doc
	if err != nil {
		print(err.Error())

	}
	var result []Doc
	for _, item := range put3.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Doc)
		fmt.Printf("%#v\n", t)
		result = append(result, t)
	}
	return result, err
}

func DeleteFromES(docId string) { //指定想要删除的文档的docId
	client, _ := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://elasticsearch.DianasDog.secoder.local:9200"))
	var err error
	res, err := client.Delete().Index("document").
		Id(docId).
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

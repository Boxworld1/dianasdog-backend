package elastic_search

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

func Insert(docId string, content string) {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
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
		return
	}
	fmt.Printf("Indexed DocId %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}

func Update(docId string, newContent string) {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	//下面是更新项目的函数，需要传入docid，对数据进行改变
	put2, err := client.Update().
		Index("document").
		Id(docId).
		Doc(map[string]interface{}{"content": newContent}).
		Do(context.Background())
	if err != nil {
		print(err.Error())
		return
	}
	fmt.Printf("update content %s\n", put2.Result)

}

func Search(content string) { //按照内容去查找，不是精确查找，只要有匹配词就可以
	client, _ := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	var typ Doc
	var err error
	var put4 *elastic.SearchResult
	matchPhraseQuery := elastic.NewMatchPhraseQuery("content", content)
	put4, err = client.Search("document").Query(matchPhraseQuery).Do(context.Background())
	if err != nil {
		print(err.Error())
		return
	}
	for _, item := range put4.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法，一条条把数据列出来
		t := item.(Doc)
		fmt.Printf("%#v\n", t)
	}
}

func FetchAll() { //拿到类型document里的所有数据
	client, _ := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	var put3 *elastic.SearchResult
	var err error
	//取所有
	put3, err = client.Search("document").Do(context.Background())
	var typ Doc
	if err != nil {
		print(err.Error())
		return
	}
	for _, item := range put3.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Doc)
		fmt.Printf("%#v\n", t)
	}

}

func Delete(docId string) { //指定想要删除的文档的docId
	client, _ := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
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

func main() {
	Insert("1", "宝马760")

	//下面是更新项目的函数，需要传入docid，对数据进行改变

	Update("1", "奔驰e系")

	//	下面是搜索功能
	FetchAll()
	//删除一条数据,只需要它的index和id
	//Delete("1")
	//搜索
	Search("桑塔纳")
}

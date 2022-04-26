// @title	TestRouter
// @description	此函数的用途为检查 SetupRouter 函数的正确性
// @auth	ryl		2022/4/13	18:00
// @param	t		*testing.T	testing 用参数

package communication

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

type MapStruct struct {
	key   string
	value string
}

func TestRouter(t *testing.T) {
	// 定义测试用例
	// 分別记录了正确的返回码和内容
	tests := []struct {
		result []int
		param  []MapStruct
	}{
		{[]int{200, 400, 400, 400, 400}, []MapStruct{
			{"content", `{"query": "apple"}`},
		}},
		{[]int{400, 400, 400, 400, 200}, []MapStruct{
			{"content", `{"username": "tester"}`},
		}},
		{[]int{400, 200, 400, 400, 400}, []MapStruct{
			{"resource", "testcase_car"},
			{"data", "{\"resource\":\"testcase_car\",\"write_setting\":{\"key\":{\"dump_digest\":\"true\",\"dump_invert_idx\":\"true\",\"dump_dict\":\"true\"},\"display.title\":{\"dump_digest\":\"false\",\"dump_invert_idx\":\"false\",\"dump_dict\":\"true\"}}}"},
		}},
		{[]int{400, 400, 200, 400, 400}, []MapStruct{
			{"resource", "testcase_car"},
			{"data", "{\"resource\":\"testcase_car\",\"rule_recall_setting_list\":{\"garbage_dict_list\":[\"parser_\u58a8\u8ff9\u5929\u6c14_200_garbage\"],\"pattern_list\":[{\"pattern_item_array\":[[{\"data_pointer\":\"name\",\"data_source_type\":0,\"field_name\":\"name\",\"query_item_type\":0}],[{\"data_pointer\":\"parser_\u58a8\u8ff9\u5929\u6c14_200_intent\",\"data_source_type\":4,\"field_name\":\"intent\",\"query_item_type\":2}]],\"use_common_garbage_dict\":true,\"use_synonym_dict\":true}],\"pre_processors\":[],\"synonym_dict_list\":[\"parser_synonym_\u58a8\u8ff9\u5929\u6c14_200_1622687488\"]}}"},
		}},
		{[]int{400, 400, 400, 200, 400}, []MapStruct{
			{"type", "insert"},
			{"resource", "testcase_car"},
			{"filename", "testcase_car.xml"},
			{"data", "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?> <DOCUMENT> <item> <key>红豆词1</key> <display> <title>红豆词</title> <score>425</score> <name0>红豆词</name0> <writer>王国维</writer> <dynasty>近代</dynasty> <detail><![CDATA[一南国秋深可奈何，手持红豆几摩挲。累累本是无情物，谁把闲愁付与他。	二门外青骢郭外舟，人生无奈是离愁。不辞苦向东风祝，到处人间作石尤。	三别浦盈盈水又波，凭栏渺渺思如何？纵教踏破江南种，只恐春来茁更多。	四匀圆万颗争相似，暗数千回不厌痴。留取他年银烛下，拈来细与话相思。]]></detail> <url><![CDATA[http://m.tool.liuxue86.com/shici_view_9b9c4a43ac9b9c4a/?ref=tt]]></url> <content><![CDATA[一南国秋深可奈何]]></content> <content><![CDATA[手持红豆几摩挲]]></content> <content><![CDATA[累累本是无情物]]></content> <content><![CDATA[谁把闲愁付与他]]></content> <content><![CDATA[二门外青骢郭外舟]]></content> <content><![CDATA[人生无奈是离愁]]></content> <content><![CDATA[不辞苦向东风祝]]></content> <content><![CDATA[到处人间作石尤]]></content> <content><![CDATA[三别浦盈盈水又波]]></content> <content><![CDATA[凭栏渺渺思如何]]></content> <content><![CDATA[纵教踏破江南种]]></content> <content><![CDATA[只恐春来茁更多]]></content> <content><![CDATA[四匀圆万颗争相似]]></content> <content><![CDATA[暗数千回不厌痴]]></content> <content><![CDATA[留取他年银烛下]]></content> <content><![CDATA[拈来细与话相思]]></content> </display> </item> <item> <key>读史二十首2</key> <display> <title>读史二十首</title> <score>323</score> <name0>读史二十首</name0> <writer>王国维</writer> <dynasty>近代</dynasty> <detail><![CDATA[	一回首西陲势渺茫，东迁种族几星霜?何当踏破双芒屐，却向昆仑望故乡。	二两条云岭摩天出，九曲黄河绕地回。自是当年游牧地，有人曾号伏羲来。	三及及生存起竞争，流传神话使人惊。铜头铁额今安在？始信轩皇苦用兵。	四澶漫江淮万里春，九黎才格又苗民。即今腿髻穷山里，此是江南旧主人。	五二帝精魂死不孤，嵇山陵庙似苍梧。耄年未罢征苗旅，神武如斯旷代无。	六铜刀岁岁战东欧，石弩年年出挹娄。毕竟中原开化早，已闻昉铁贡梁州。	七谁向钧天听乐过，秦中自古鬼神多。即今《诅楚文》犹在，才告巫咸又亚驼。	八《春秋》谜语苦难诠，历史开山数腐迁。前后固应无此文，一书上下两千年。	九汉作昆池始见煤，当年赀力信雄哉。于今莫笑胡僧妄，本是洪荒劫后灰。	十挥戈大启汉山河，武帝雄才世讵多。轻骑今朝绝大漠，楼川明日下洋河。	十一惠光东照日炎炎，河陇降王正款边。不是金人先入汉，永平谁证梦中缘。	十二西域纵横尽百城，张陈远略逊甘英。千秋壮观君知否？黑海东头望大秦。	十三三方并帝古未有，两贤向厄我所闻。何来洒落樽前语：天下英雄惟使君。	十四北临洛水拜陵园，奉表迁都大义存。纵使暮年终作贼，江东那更有桓温。	十五江南天子皆词客，河北诸王尽将才。乍歌乐府《兰陵曲》，又见湘东玉轴灰。	十六晋阳蜿蜿起飞龙，北面倾心事犬戎。亲出渭桥擒诘利，文皇端不愧英雄。	十七南海商船来大食，西京袄寺建波斯。远人尽有如归乐，知是唐家全盛时。	十八五国风光惨不支，崖山波浪浩无牙。当年国势凌迟甚，争怪诸贤唱攘夷。	十九黑水金山启伯图，长驱远摭世间无。至今碧眼黄须客，犹自惊魂说拔都。	二十东海人奴盖世雄，卷舒八道势如风。碧蹄倘得擒渠反，大壑何由起蜇龙。]]></detail> <url><![CDATA[http://m.tool.liuxue86.com/shici_view_9b9c4b43ac9b9c4b/?ref=tt]]></url> <content><![CDATA[一回首西陲势渺茫]]></content> <content><![CDATA[东迁种族几星霜]]></content> <content><![CDATA[何当踏破双芒屐]]></content> <content><![CDATA[却向昆仑望故乡]]></content> <content><![CDATA[二两条云岭摩天出]]></content> <content><![CDATA[九曲黄河绕地回]]></content> <content><![CDATA[自是当年游牧地]]></content> <content><![CDATA[有人曾号伏羲来]]></content> <content><![CDATA[三及及生存起竞争]]></content> <content><![CDATA[流传神话使人惊]]></content> <content><![CDATA[铜头铁额今安在]]></content> <content><![CDATA[始信轩皇苦用兵]]></content> <content><![CDATA[四澶漫江淮万里春]]></content> <content><![CDATA[九黎才格又苗民]]></content> <content><![CDATA[即今腿髻穷山里]]></content> <content><![CDATA[此是江南旧主人]]></content> <content><![CDATA[五二帝精魂死不孤]]></content> <content><![CDATA[嵇山陵庙似苍梧]]></content> <content><![CDATA[耄年未罢征苗旅]]></content> <content><![CDATA[神武如斯旷代无]]></content> <content><![CDATA[六铜刀岁岁战东欧]]></content> <content><![CDATA[石弩年年出挹娄]]></content> <content><![CDATA[毕竟中原开化早]]></content> <content><![CDATA[已闻昉铁贡梁州]]></content> <content><![CDATA[七谁向钧天听乐过]]></content> <content><![CDATA[秦中自古鬼神多]]></content> <content><![CDATA[即今《诅楚文》犹在]]></content> <content><![CDATA[才告巫咸又亚驼]]></content> <content><![CDATA[八《春秋》谜语苦难诠]]></content> <content><![CDATA[历史开山数腐迁]]></content> <content><![CDATA[前后固应无此文]]></content> <content><![CDATA[一书上下两千年]]></content> <content><![CDATA[九汉作昆池始见煤]]></content> <content><![CDATA[当年赀力信雄哉]]></content> <content><![CDATA[于今莫笑胡僧妄]]></content> <content><![CDATA[本是洪荒劫后灰]]></content> <content><![CDATA[十挥戈大启汉山河]]></content> <content><![CDATA[武帝雄才世讵多]]></content> <content><![CDATA[轻骑今朝绝大漠]]></content> <content><![CDATA[楼川明日下洋河]]></content> <content><![CDATA[十一惠光东照日炎炎]]></content> <content><![CDATA[河陇降王正款边]]></content> <content><![CDATA[不是金人先入汉]]></content> <content><![CDATA[永平谁证梦中缘]]></content> <content><![CDATA[十二西域纵横尽百城]]></content> <content><![CDATA[张陈远略逊甘英]]></content> <content><![CDATA[千秋壮观君知否]]></content> <content><![CDATA[黑海东头望大秦]]></content> <content><![CDATA[十三三方并帝古未有]]></content> <content><![CDATA[两贤向厄我所闻]]></content> <content><![CDATA[何来洒落樽前语]]></content> <content><![CDATA[天下英雄惟使君]]></content> <content><![CDATA[十四北临洛水拜陵园]]></content> <content><![CDATA[奉表迁都大义存]]></content> <content><![CDATA[纵使暮年终作贼]]></content> <content><![CDATA[江东那更有桓温]]></content> <content><![CDATA[十五江南天子皆词客]]></content> <content><![CDATA[河北诸王尽将才]]></content> <content><![CDATA[乍歌乐府《兰陵曲》]]></content> <content><![CDATA[又见湘东玉轴灰]]></content> <content><![CDATA[十六晋阳蜿蜿起飞龙]]></content> <content><![CDATA[北面倾心事犬戎]]></content> <content><![CDATA[亲出渭桥擒诘利]]></content> <content><![CDATA[文皇端不愧英雄]]></content> <content><![CDATA[十七南海商船来大食]]></content> <content><![CDATA[西京袄寺建波斯]]></content> <content><![CDATA[远人尽有如归乐]]></content> <content><![CDATA[知是唐家全盛时]]></content> <content><![CDATA[十八五国风光惨不支]]></content> <content><![CDATA[崖山波浪浩无牙]]></content> <content><![CDATA[当年国势凌迟甚]]></content> <content><![CDATA[争怪诸贤唱攘夷]]></content> <content><![CDATA[十九黑水金山启伯图]]></content> <content><![CDATA[长驱远摭世间无]]></content> <content><![CDATA[至今碧眼黄须客]]></content> <content><![CDATA[犹自惊魂说拔都]]></content> <content><![CDATA[二十东海人奴盖世雄]]></content> <content><![CDATA[卷舒八道势如风]]></content> <content><![CDATA[碧蹄倘得擒渠反]]></content> <content><![CDATA[大壑何由起蜇龙]]></content> </display> </item> </DOCUMENT>"},
		}},
		{[]int{400, 400, 400, 200, 400}, []MapStruct{
			{"type", "delete"},
			{"resource", "testcase_car"},
			{"filename", "testcase_car.xml"},
			{"data", "<DOCUMENT> </DOCUMENT>"},
		}},
		{[]int{400, 400, 400, 200, 400}, []MapStruct{
			{"type", "update"},
			{"resource", "testcase_car"},
			{"filename", "testcase_car.xml"},
			{"data", "<DOCUMENT> </DOCUMENT>"},
		}},
		{[]int{400, 400, 400, 400, 400}, []MapStruct{}},
		{[]int{400, 400, 400, 400, 400}, []MapStruct{
			{"resource", "testcase_car"},
		}},
		{[]int{400, 400, 400, 400, 400}, []MapStruct{
			{"type", "update"},
			{"resource", "testcase_car"},
		}},
	}

	// 定义要测试的接口
	methods := []MapStruct{
		{"POST", "/search"},
		{"POST", "/setting"},
		{"POST", "/pattern"},
		{"POST", "/data"},
		{"POST", "/login"},
	}

	// 开启 router
	router := SetupRouter()

	for key, method := range methods {
		for dataID, testcase := range tests {
			var req *http.Request

			// mock 一个 HTTP 请求
			if dataID <= 1 {
				// json raw data for /search
				req = httptest.NewRequest(
					method.key,   // 请求方法
					method.value, // 请求 URL
					strings.NewReader(testcase.param[0].value), // 请求参数
				)
			} else {
				// form data for other ports
				// 根据测试用例加入参数
				form := url.Values{}
				for _, value := range testcase.param {
					key := value.key
					content := value.value
					form.Add(key, content)
				}

				req = httptest.NewRequest(
					method.key,                       // 请求方法
					method.value,                     // 请求 URL
					strings.NewReader(form.Encode()), // 请求参数
				)

				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
			}

			// mock 一个响应记录器
			w := httptest.NewRecorder()

			// 让 server 端处理 mock 请求并记录返回的响应内容
			router.ServeHTTP(w, req)

			// 校验状态码是否符合预期
			if testcase.result[key] != w.Code {
				fmt.Println("testcase:", key, "with data:", dataID, "get:", w.Code, "but expect:", testcase.result[key])
				t.Error("状态码错误")
			}
		}
	}
}

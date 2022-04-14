// @title	TestRouter
// @description	此函数的用途为检查 SetupRouter 函数的正确性
// @auth	ryl		2022/4/13	18:00
// @param	t		*testing.T	testing 用参数

package communication

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouter(t *testing.T) {
	// 定义测试用例
	// 分別记录了正确的返回码和内容
	tests := []struct {
		result []int
		param  string
	}{
		{[]int{200, 400, 400, 400}, `{
			"query": "apple"
		}`},
		{[]int{400, 200, 400, 400}, `{
			"resource": "testcase_car",
			"write_setting": {
				"a.b.c": {
					"dump_digest": "true",
					"dump_invert_idx": "false",
					"dump_dict": "true"
				},
				"a.e": {
					"dump_digest": "false",
					"dump_invert_idx": "false",
					"dump_dict": "true"
				},
				"b.g": {
					"dump_digest": "true",
					"dump_invert_idx": "true",
					"dump_dict": "true"
				},
				"f.a": {
					"dump_digest": "true",
					"dump_invert_idx": "false",
					"dump_dict": "true"
				}
			}
		}`},
		{[]int{400, 400, 200, 400}, `{
			"type": "insert", 
			"resource": "testcase_car",
			"file": "testcase_car.xml",
			"data": " "
		}`},
		{[]int{400, 400, 200, 400}, `{
			"type": "delete", 
			"resource": "testcase_car",
			"file": "testcase_car.xml",
			"data": " "
		}`},
		{[]int{400, 400, 200, 400}, `{
			"type": "update", 
			"resource": "testcase_car",
			"file": "testcase_car.xml",
			"data": " "
		}`},
		{[]int{400, 400, 400, 200}, `{
			"resource": "testcase_car", 
			"data": {
				
			}
		}`},
		{[]int{400, 400, 400, 400}, "{}"},
	}

	// 定义要调用的接口
	methods := []struct {
		method string
		url    string
	}{
		{"POST", "/search"},
		{"POST", "/setting"},
		{"POST", "/data"},
		{"POST", "/pattern"},
	}
	router := SetupRouter()

	for key, method := range methods {
		for _, testcase := range tests {
			// mock一个HTTP请求
			req := httptest.NewRequest(
				method.method,                     // 请求方法
				method.url,                        // 请求URL
				strings.NewReader(testcase.param), // 请求参数
			)

			// mock一个响应记录器
			w := httptest.NewRecorder()

			// 让server端处理mock请求并记录返回的响应内容
			router.ServeHTTP(w, req)

			// 校验状态码是否符合预期
			if testcase.result[key] != w.Code {
				fmt.Println(key, w.Code)
				t.Error("状态码错误")
			}
		}
	}
}

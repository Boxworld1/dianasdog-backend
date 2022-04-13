// @title	SetupRouter
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

func Test_helloHandler(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		id    int
		param string
	}{
		{0, `{"query": "apple"}`},
		{1, `{
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
		{2, `{"type": "insert", "resource":"car"}`},
		{3, `{"resource": "car", "data":""}`},
		{4, "{}"},
	}

	// 定义接口
	methods := []struct {
		method string
		url    string
	}{
		{"POST", "/search"},
		{"POST", "/setting"},
		// {"POST", "/data"},
		// {"POST", "/pattern"},
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
			if (testcase.id == key && w.Code != 200) || (testcase.id != key && w.Code != 400) {
				fmt.Println(testcase.id, w.Code)
				t.Error("状态码错误")
			}
		}
	}
}

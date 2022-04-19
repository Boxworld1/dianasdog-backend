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
		{[]int{200, 400, 400, 400}, []MapStruct{
			{"content", `{"query": "apple"}`},
		}},
		{[]int{400, 200, 400, 400}, []MapStruct{
			{"resource", "testcase_car"},
			{"data", "{\"resource\":\"testcase_car\",\"write_setting\":{\"a.b.c\":{\"dump_digest\":\"true\",\"dump_invert_idx\":\"false\",\"dump_dict\":\"true\"},\"a.e\":{\"dump_digest\":\"false\",\"dump_invert_idx\":\"false\",\"dump_dict\":\"true\"},\"b.g\":{\"dump_digest\":\"true\",\"dump_invert_idx\":\"true\",\"dump_dict\":\"true\"},\"f.a\":{\"dump_digest\":\"true\",\"dump_invert_idx\":\"false\",\"dump_dict\":\"true\"}}}"},
		}},
		{[]int{400, 400, 200, 400}, []MapStruct{
			{"type", "insert"},
			{"resource", "testcase_car"},
			{"file", "testcase_car.xml"},
			{"data", "no"},
		}},
		{[]int{400, 400, 200, 400}, []MapStruct{
			{"type", "delete"},
			{"resource", "testcase_car"},
			{"file", "testcase_car.xml"},
			{"data", "nod"},
		}},
		{[]int{400, 400, 200, 400}, []MapStruct{
			{"type", "update"},
			{"resource", "testcase_car"},
			{"file", "testcase_car.xml"},
			{"data", "nod"},
		}},
		{[]int{400, 400, 400, 400}, []MapStruct{
			{"resource", "testcase_car"},
			{"data", "nods"},
		}},
		{[]int{400, 400, 400, 400}, []MapStruct{}},
	}

	// 定义要调用的接口
	methods := []MapStruct{
		{"POST", "/search"},
		// {"POST", "/setting"},
		// {"POST", "/data"},
		// {"POST", "/pattern"},
	}

	// 开启 router
	router := SetupRouter()

	for key, method := range methods {
		for dataID, testcase := range tests {
			var req *http.Request

			// mock 一个 HTTP 请求
			if dataID == 0 {
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
					form.Add(value.key, value.value)
				}
				req = httptest.NewRequest(
					method.key,                       // 请求方法
					method.value,                     // 请求 URL
					strings.NewReader(form.Encode()), // 请求参数
				)
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

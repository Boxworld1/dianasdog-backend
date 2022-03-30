// @title		SetupRouter
// @description	此函数的用途为检查 SetupRouter 函数的正确性
// @auth		ryl				2022/3/30		22:00
// @param		t				*testing.T		testing 用参数

package communication

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_helloHandler(t *testing.T) {
	// 定义两个测试用例
	tests := []struct {
		id    int
		param string
	}{
		{1, `{"query": "apple"}`},
		{2, "{}"},
	}

	router := SetupRouter()

	for _, testcase := range tests {
		// mock一个HTTP请求
		req := httptest.NewRequest(
			"GET",                             // 请求方法
			"/search",                         // 请求URL
			strings.NewReader(testcase.param), // 请求参数
		)

		// mock一个响应记录器
		w := httptest.NewRecorder()

		// 让server端处理mock请求并记录返回的响应内容
		router.ServeHTTP(w, req)

		// 校验状态码是否符合预期
		if (testcase.id == 1 && w.Code != 200) || (testcase.id == 2 && w.Code != 400) {
			fmt.Println(testcase.id, w.Code)
			t.Error("状态码错误")
		}
	}
}

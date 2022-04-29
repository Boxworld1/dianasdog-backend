// @title	TestAhocorasick
// @description	检查 AC 自动机
// @auth	jz	2022/4/7	12:00
// @param	t	*testing.T	testing 用参数

package search

import (
	"testing"
)

// 测试 AC 自动机的匹配功能
func TestMatch(t *testing.T) {
	ac := NewMatcher()
	dict := []string{"宝马", "奔驰", "奥迪"}
	ac.Build(dict)
	ret := ac.Match("宝马的价格是多少")
	if ret[0].BegPosition != 0 || ret[0].EndPosition != 1 {
		t.Error("wrong answer")
	}
}

// 测试 AC 自动机的返回值
func TestCheck(t *testing.T) {
	ac := NewMatcher()
	dict := []string{"she", "her", "he", "say"}
	ac.Build(dict)
	ret1 := ac.Check("shershx")
	ret2 := ac.Check("shfjk")
	if ret1 != true || ret2 != false {
		t.Error("wrong answer")
	}
}

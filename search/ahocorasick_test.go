package search

import (
	"testing"
)

func TestMatch(t *testing.T) {
	ac := NewMatcher()
	dict := []string{"she", "her", "he", "say"}
	ac.Build(dict)
	ret := ac.Match("shershx")
	if ret[0].BegPosition != 0 || ret[0].EndPosition != 2 || ret[1].BegPosition != 1 || ret[1].EndPosition != 2 || ret[2].BegPosition != 1 || ret[2].EndPosition != 3 {
		t.Error("wrong answer")
	}

}

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

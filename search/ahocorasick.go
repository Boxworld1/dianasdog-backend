// @title	Ahocorasick
// @description	AC 自动机
// @auth	jz	2022/4/7	12:00

package search

import (
	"container/list"
)

type trieNode struct {
	count int
	fail  *trieNode
	child map[rune]*trieNode
	len   int
}

// @title: newTrieNode
// @description:create a new TrieNode
// @param: do not need a param
// @return: *trieNode    return a pointer of the TrieNode
func newTrieNode() *trieNode {
	return &trieNode{
		count: 0,
		fail:  nil,
		child: make(map[rune]*trieNode),
		len:   0,
	}
}

type Matcher struct {
	root *trieNode
	size int
}

type Term struct {
	BegPosition int
	EndPosition int
}

// @title: NewMatcher
// @description:create a new matcher
// @param: do not need a param
// @return: *Matcher    return a pointer of the new matcher
func NewMatcher() *Matcher {
	return &Matcher{
		root: newTrieNode(),
		size: 0,
	}
}

// @title: BuildNewMatcher
// @description:build a newmatcher from the dictionary
// @param: dictionary    []string    the dict used to build the matcher
// @return: *Matcher     return a pointer of the new matcher
func BuildNewMatcher(dictionary []string) *Matcher {
	m := &Matcher{
		root: newTrieNode(),
		size: 0,
	}
	m.Build(dictionary)
	return m
}

// @title: Build
// @description:build the matcher from the dictionary
// @param: dictionary    []string    the dict used to build the matcher
// @return: *Matcher     return a pointer of the new matcher
func (m *Matcher) Build(dictionary []string) {
	for i := range dictionary {
		m.insert(dictionary[i])
	}
	m.buildfail()
}

// @title: Match
// @description:string match search
// @param: s    string    the string needed to be searched
// @return: []*Term       return all templates matched as their positions on targeted string
func (m *Matcher) Match(s string) []*Term {
	curNode := m.root
	var p *trieNode = nil

	//	mark := make([]bool, m.size)
	ret := make([]*Term, 0)

	for index, rune := range []rune(s) {
		for curNode.child[rune] == nil && curNode != m.root {
			curNode = curNode.fail
		}
		curNode = curNode.child[rune]
		if curNode == nil {
			curNode = m.root
		}

		p = curNode
		for p != m.root && p.count > 0 {
			for i := 0; i < p.count; i++ {
				ret = append(ret, &Term{BegPosition: index - p.len + 1, EndPosition: index})
			}
			p = p.fail
		}
	}

	return ret
}

// @title: check
// @description:check  whether s can match any template in ac
// @param: s    string    the string needed to be checked
// @return: bool    return the result of the check
func (m *Matcher) Check(s string) bool {
	curNode := m.root
	var p *trieNode = nil
	for _, rune := range s {
		for curNode.child[rune] == nil && curNode != m.root {
			curNode = curNode.fail
		}
		curNode = curNode.child[rune]
		if curNode == nil {
			curNode = m.root
		}

		p = curNode
		if p != m.root && p.count > 0 {
			return true
		}
	}
	return false
}

// @title: build
// @description:initialize the fail of the ac
// @param: do not need a param
// @return: do not need a return-value
func (m *Matcher) buildfail() {
	ll := list.New()
	ll.PushBack(m.root)
	for ll.Len() > 0 {
		temp := ll.Remove(ll.Front()).(*trieNode)
		var p *trieNode = nil

		for i, v := range temp.child {
			if temp == m.root {
				v.fail = m.root
			} else {
				p = temp.fail
				for p != nil {
					if p.child[i] != nil {
						v.fail = p.child[i]
						break
					}
					p = p.fail
				}
				if p == nil {
					v.fail = m.root
				}
			}
			ll.PushBack(v)
		}
	}
}

// @title: insert
// @description:insert a string to the ac
// @param: s    string    the string needed to be inserted
// @return: do not need a return-value
func (m *Matcher) insert(s string) {
	curNode := m.root
	for _, v := range s {
		if curNode.child[v] == nil {
			curNode.child[v] = newTrieNode()
		}
		curNode = curNode.child[v]
	}
	curNode.count++
	curNode.len = len([]rune(s))
	m.size++
}

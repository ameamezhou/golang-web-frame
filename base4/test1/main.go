package main

import (
	"fmt"
	"strings"
)

// Tire desc
type node struct {
	pattern     string   // 完整的路由        /xiao/qi/:zhou/test
	part        string   // 路由中的一部分     :zhou
	children 	[]*node  // 存储子节点
	isWild  	bool     // 是否精确匹配       在含有: or * 的时候为 true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 根据理解重新写的个人逻辑
//func (t *node) matchChild(part string) *node {
//
//	// 先搜寻精确匹配路由
//	for _, child := range t.children{
//		if child.part == part && !child.isWild {
//			return child
//		}
//	}
//	// 再搜寻模糊匹配路由
//	for _, child := range t.children{
//		if child.part == part && child.isWild {
//			return child
//		}
//	}
//
//	return nil
//}

func (n *node) insert(pattern string, parts []string, height int) {
	//fmt.Println(n)
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	//fmt.Println(child == nil)
	if child == nil {
		//fmt.Println(1)
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
		//fmt.Println(n)
		//fmt.Println(child)
	}
	child.insert(pattern, parts, height+1)
}

func main(){
	a := node{
		pattern: "/",
		part: "",
		children: []*node{},
		isWild: false,
	}
	pathSlice := []string{
		"/xiao/qi/:xia",
		"/xiao/qi/zhou",
		"/xiao/qi/*",
	}
	for _, v := range pathSlice {
		parts := strings.Split(v, "/")
		//fmt.Println(parts)
		//fmt.Println(v)
		a.insert(v, parts, 1)
	}
	fmt.Println(a)
	for _, v1 := range a.children {
		fmt.Println(*v1)
		for _, v2 := range (*v1).children {
			fmt.Println(*v2)
			for _, v3 := range (*v2).children {
				fmt.Println(*v3)
			}
		}
	}

}

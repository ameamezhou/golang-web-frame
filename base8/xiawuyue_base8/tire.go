package xiawuyue_base8

import (
	"fmt"
	"strings"
)

// 首先设计树节点上应该存储的信息

// Tire desc
type Tire struct {
	pattern     string   // 完整的路由        /xiao/qi/:zhou/test
	part        string   // 路由中的一部分     :zhou
	children 	[]*Tire  // 存储子节点
	isWild  	bool     // 是否精确匹配       在含有: or * 的时候为 true
}

// 第一个匹配成功的节点，用于插入
func (t *Tire) matchChild(part string) *Tire {
	for _, child := range t.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}
// 所有匹配成功的节点，用于查找
func (t *Tire) matchChildren(part string) []*Tire {
	nodes := make([]*Tire, 0)
	for _, child := range t.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert desc
func (t *Tire) insert(pattern string, parts []string, depth int) error {
	if len(parts) == depth {
		if t.pattern != "" {
			if t.pattern == pattern {
				return fmt.Errorf("pattern: %s is exist. please check and update ~", pattern)
			}
		}
		t.pattern = pattern
		return nil
	}

	part := parts[depth]
	child := t.matchChild(part)
	if child == nil {
		// 当发现没有匹配到对应part的子节点的时候  就要新建一个子节点
		child = &Tire{part: part, isWild: part[0] == ':' || part[0] == '*'}
		t.children = append(t.children, child)
	}
	err := child.insert(pattern, parts, depth + 1)
	return err
}

func (t *Tire) search(parts []string, height int) *Tire {
	if len(parts) == height || strings.HasPrefix(t.part, "*") {
		if t.pattern == "" {
			return nil
		}
		return t
	}

	part := parts[height]
	children := t.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
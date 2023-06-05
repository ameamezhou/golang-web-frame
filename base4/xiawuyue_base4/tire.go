package xiawuyue_base3

import (
	"fmt"
	"strings"
)

// 简单描述下动态路由的背景
// 我们用 map[string]HandleFunc 的形式虽然很方便
// 检索效率也高，但是有一个致命的缺点：
// 当我们的路由出现 /xiao/qi/:zhou/test or /xia/wu/yue/*
// 这种动态路由时，它将不再适用，这也是我们使用动态路由的原因
// 了解了背景之后我们来开发整个动态路由的过程

/*
动态路由有很多种实现方式，支持的规则、性能等有很大的差异。
例如开源的路由实现gorouter支持在路由规则中嵌入正则表达式，例如/p/[0-9A-Za-z]+，即路径中的参数仅匹配数字和字母；
另一个开源实现httprouter就不支持正则表达式。著名的Web开源框架gin 在早期的版本，并没有实现自己的路由，
而是直接使用了httprouter，后来不知道什么原因，放弃了httprouter，自己实现了一个版本
*/

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

// insert 要实现的功能
// 1、 插入子节点，如果路径上没有相应的子节点我们就要一起插入
// 2、 在最后讲这个节点插入的时候要更新它的 pattern
// 3、 当插入到最后发现插入的这个 pattern 已经存在了的时候，我们要及时告知并返回报错  防止用户有两个逻辑要实现，最后写成了一个url
// 嘻嘻  3 就是我们自己的创新点，感谢 geektutu 和 b站up的视频

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
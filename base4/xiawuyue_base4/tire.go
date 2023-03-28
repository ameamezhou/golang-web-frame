package xiawuyue_base3

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

// 这里我们添加的 isWild 这个参数 用在模糊匹配中，当我们匹配 /xiao/qi/xiawuyue/test 这个url时
// xiao 和 qi 都能在 Tire 的子节点中精准匹配到

// 首先我们要先写一个匹配单个节点的工具函数，匹配单个节点的目的是 当我们插入一个 /xiao/qi/zhou or /xiao/qi/:zhou
// 这样的路径时，我们需要精确匹配到根路径下的某一个子节点，并且在 /qi 这个子节点中给它加入 children 中

 /*
作者原逻辑
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

这里我对这个逻辑进行了一些修改
我这里举例三种例子：
/xiao/qi/zhou
/xiao/qi/:xia
/xiao/qi/ *   （这里和go的注释冲突了  所以加了个空格）
当我在这里需要插入一条/xiao/qi/zhou/wu/yue 这样一条路径时
我在 /qi 的字节点下面匹配到 /:xia or / * 的时候 就会获取到这个子节点并且往下传递， 就等于说我的 /wu /yue 这两个子节点就会加到 /:xia
或者 / * 的下面了  这是我们不希望看到的情况，所以为了避免这个情况  这里我把逻辑进行了修改
作者的原思想我把它汇总成代码写到了 base4/test1 写了一个测试用例，作者的动态路由存在吞噬路由的可能性


*/

func startWith(part string) bool {
	if part[0] == 58{
		return true
	}
	return false
}


func (t *Tire) MatchChild(part string) *Tire {

	for _, child := range t.children{
		if child.part == part && !child.isWild {
			return child
		}
	}

	for _, child := range t.children{
		if child.part == part && child.isWild {
			return child
		}
	}

	return nil
}


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


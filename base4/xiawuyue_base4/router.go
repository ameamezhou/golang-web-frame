package xiawuyue_base3

import (
	"strings"
)

// 将router 从 xiawuyue 主体里抽离出来
// roots key eg, roots['GET'] roots['POST']
// roots 记录的是不同方法对应的前缀树路由  POST GET 方法
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']
// handlers 记录 方法 - 路径   来标注处理函数
type router struct {
	roots	 map[string]*Tire
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		roots:		make(map[string]*Tire),
		handlers: 	make(map[string]HandleFunc),
	}
}

// 这里如果要再添加路由的话我们就要考虑前缀树路由的插入了
// 所以这里我们要先将拿到的pattern进行预处理，方便我们Tire 进行insert
func parsePattern(pattern string)[]string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router)addRoute(method string, pattern string, handleFunc HandleFunc) {
	// 其中method 是用来区分 get post 等方法的
	// patter 是提到的 muxEntry 中的匹配字符串 也就是具体的路径
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &Tire{}
	}
	err := r.roots[method].insert(pattern, parts, 0)
	if err != nil {
		// log err     后续实现log代码
		return
	}
	r.handlers[key] = handleFunc
}

// 这里get route 是要用来初始化的   见handle
func (r *router)getRoute(method string, path string)(*Tire, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	// searchParts 找到根节点   然后判断最后记录 pattern节点的内容。
	// 这里的逻辑是： 找到对应的前缀树节点后，要找到通配符所匹配的字符串  也就是我们说的传进来的参数
	// 所以新建params 存储参数内容，: 开头的字符串是用来匹配某一个字符     然后*的话会匹配后面所有的字符内容
	t := root.search(searchParts, 0)
	if t != nil {
		parts := parsePattern(t.pattern)
		for i, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[i]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[i:], "/")
				break
			}
		}
		return t, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	t, params := r.getRoute(c.Method, c.Path)
	if t != nil {
		c.Params = params
		key := c.Method + "-" + t.pattern
		if handler, ok := r.handlers[key]; ok {
			handler(c)
		}
	}


}
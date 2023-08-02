package xiawuyue_base3

import "strings"

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
func (r *router) parsePattern(pattern string)[]string {
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
	parts := r.parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &Tire{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handleFunc
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	}

}
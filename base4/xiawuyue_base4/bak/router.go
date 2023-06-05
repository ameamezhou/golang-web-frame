package bak

// 将router 从 xiawuyue 主体里抽离出来

type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandleFunc)}
}

func (r *router)addRoute(method string, pattern string, handleFunc HandleFunc) {
	// 其中method 是用来区分 get post 等方法的
	// patter 是提到的 muxEntry 中的匹配字符串 也就是具体的路径
	key := method + "-" + pattern
	r.handlers[key] = handleFunc
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	}

}
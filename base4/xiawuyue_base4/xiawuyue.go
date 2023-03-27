package xiawuyue_base3

import (
	"net/http"
)

type XiaWuYue struct {
	router *router
}

// New 直接调用New方法构建对象
func New() *XiaWuYue {
	return &XiaWuYue{ router: newRouter() }
}

// HandleFunc 简单定义一类函数  这就是后续具体的处理方法的类型
type HandleFunc func(c *Context)

func (x *XiaWuYue)addRoute(method string, pattern string, handleFunc HandleFunc) {
	x.router.addRoute(method, pattern, handleFunc)
}

func (x *XiaWuYue) Get(pattern string, handleFunc HandleFunc) {
	x.addRoute("GET", pattern, handleFunc)
}

func (x *XiaWuYue) Post(pattern string, handleFunc HandleFunc) {
	x.addRoute("POST", pattern, handleFunc)
}

func (x *XiaWuYue) Pull(pattern string, handleFunc HandleFunc) {
	x.addRoute("PULL", pattern, handleFunc)
}

func (x *XiaWuYue) Delete(pattern string, handleFunc HandleFunc) {
	x.addRoute("DELETE", pattern, handleFunc)
}

// Run 这里将run函数独立出来，后面我们就不用再使用http包进行跑服务了  直接用xiawuyue.run就好了
func (x *XiaWuYue) Run(addr string) {
	http.ListenAndServe(addr, x)
}

// 这里因为我们新建了context 所以我们只需要将context传给我们抽离出来的router使用就好了

func (x *XiaWuYue) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	x.router.handle(c)
}
package xiawuyue_base5

import (
	"net/http"
)

type XiaWuYue struct {
	router *router
}

/*
我们理想的调用 Group 的方式是
r := gee.New()
v1 := r.Group("/v1")
v1.GET("/", func(c *gee.Context) {
	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
})
拿到一个group 并在这个 group 中进行添加后续的路由

那么我们就要思考一个 Group 对象需要哪些属性
首先我们要能够记录前缀，因为这个要记录前缀  我们补充后面内容
我们还需要知道当前分组的父组件是谁
web框架最重要的就是中间件，所以我们还要存中间件（其实也就是对应着一个预处理的函数）
最后我们要用一个 (*Engine.addRoute())
*/

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
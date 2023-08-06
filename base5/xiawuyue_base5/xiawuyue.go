package xiawuyue_base5

import (
	"log"
	"net/http"
)


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
最后我们要用一个 (*Engine.addRoute()) 来映射所有的路由跪着和 Handler。
*/

type RouterGroup struct {
	prefix		string
	middlewares []HandlerFunc
	parent 		*RouterGroup
	engine		*XiaWuYue			// all groups share a Engine(XiaWuYue) instance
}

// XiaWuYue 这里我们将Engine改为XiaWuYue
type  XiaWuYue struct {
	*RouterGroup
	router 	*router
	groups	[]*RouterGroup
}

// New 直接调用New方法构建对象
func New() *XiaWuYue {
	qiuWu := &XiaWuYue{ router: newRouter() }
	qiuWu.RouterGroup = &RouterGroup{ engine: qiuWu }
	qiuWu.groups = []*RouterGroup{ qiuWu.RouterGroup }

	return &XiaWuYue{ router: newRouter() }
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}

	return newGroup
}

// HandlerFunc 简单定义一类函数  这就是后续具体的处理方法的类型
type HandlerFunc func(c *Context)

func (g *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("roter %4s - %s", method, pattern)
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *RouterGroup) Get(pattern string, handlerFunc HandlerFunc) {
	g.addRoute("GET", pattern, handlerFunc)
}

func (g *RouterGroup) Post(pattern string, handleFunc HandlerFunc) {
	g.addRoute("POST", pattern, handleFunc)
}

func (g *RouterGroup) Pull(pattern string, handleFunc HandlerFunc) {
	g.addRoute("PULL", pattern, handleFunc)
}

func (g *RouterGroup) Delete(pattern string, handleFunc HandlerFunc) {
	g.addRoute("DELETE", pattern, handleFunc)
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
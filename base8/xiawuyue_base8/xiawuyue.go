package xiawuyue_base8

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

type RouterGroup struct {
	prefix		string
	middlewares []HandlerFunc
	parent 		*RouterGroup
	engine		*XiaWuYue			// all groups share a Engine(XiaWuYue) instance
}

type  XiaWuYue struct {
	*RouterGroup
	router 	*router
	groups	[]*RouterGroup
	htmlTemplates *template.Template // for html render
	funcMap       template.FuncMap   // for html render
}

// New 直接调用New方法构建对象
func New() *XiaWuYue {
	qiuWu := &XiaWuYue{ router: newRouter() }
	qiuWu.RouterGroup = &RouterGroup{ engine: qiuWu }
	qiuWu.groups = []*RouterGroup{ qiuWu.RouterGroup }

	return qiuWu
}

/*
首先为 Engine 示例添加了 *template.Template 和 template.FuncMap对象，前者将所有的模板加载进内存，后者是所有的自定义模板渲染函数。

另外，给用户分别提供了设置自定义渲染函数funcMap和加载模板的方法。
*/

func (x *XiaWuYue) SetFuncMap(funcMap template.FuncMap) {
	x.funcMap = funcMap
}

func (x *XiaWuYue) LoadHTMLGlob(pattern string) {
	x.htmlTemplates = template.Must(template.New("").Funcs(x.funcMap).ParseGlob(pattern))
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

func (g *RouterGroup) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
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

// create static handler
func (g *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(g.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// serve static files
func (g *RouterGroup) Static(relativePath string, root string) {
	handler := g.createStaticHandler(relativePath, http.Dir(root))
	/*
	之前还奇怪为什么要写
	handler := group.createStaticHandler(relativePath, http.Dir(root)) urlPattern := path.Join(relativePath, "/*filepath")
	后来才发现原来fileHandler.ServeHTTP 会把req.url.path 作为文件路径
	*/
	urlPattern := path.Join(relativePath, "/*filepath")
	g.Get(urlPattern, handler)
}

func (x *XiaWuYue) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range x.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := NewContext(w, req)
	c.handlers = middlewares
	c.xiawuyue = x
	x.router.handle(c)
}
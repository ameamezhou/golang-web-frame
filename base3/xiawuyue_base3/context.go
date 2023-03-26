package xiawuyue_base3

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// this context struct 我们需要包含两个所需变量
// 一个是我们需要写的 Writer 一个是请求 Request
// StatusCode 状态码
// Path 路径
// Method 用的方法

// Context desc
type Context struct {
	Writer		http.ResponseWriter
	Req     	*http.Request
	StatusCode 	int
	Path 		string
	Method      string
}

type Z map[string]interface{}

// 这里给 map[string]interface{} 的类型定义一个别名 我们后续在构造请求头的时候能方便一些

// NewContext 瑕无月 Context 的构造函数
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	// 这里我们需要的 Path 和 Method 都可以从 request 的参数中获取
	return &Context{
		Req: r,
		Writer: w,
		Path: r.URL.Path,
		Method: r.Method,
	}
}

// 我们在用一些框架的时候，我们发现我们前后端传输数据用到的 form 表单
// 在很多框架中我们能通过一些 key 去获取对应的value 那我们具体要怎么实现这个功能呢？
// 就需要构造这样一个函数来进行实现

// PostForm desc
func (c *Context)PostForm(key string) string {
	return c.Req.FormValue(key)
}

// 这是一个方便 我们 直接从 请求的url中获取请求参数的一些方法
// like http://xxxx.xxxx.com/xx/xxx/xzxx?abc=???&bde=???
// 获取后面的 abc 和 bde

// Query desc
func (c *Context)Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 方便我们在这个上下文中写入请求码的功能

// Status desc
func (c *Context)Status(status int) {
	c.StatusCode = status
	c.Writer.WriteHeader(status)
}

// 这是方便我们通过设置key和value去很方便的设置请求头字段的一个方法

// SetHeader desc
func (c *Context)SetHeader(key string, value string){
	c.Writer.Header().Set(key, value)
}

// 构造一些快速响应的方法

// String
func (c *Context)String(status int, format string, values ...interface{}){
	c.SetHeader("Content-Type", "text/plain")
	c.Status(status)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Json(status int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(status)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		fmt.Println("编码错误")
		// 后续开发一个 log 包
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) HTML(status int, html string) {
	c.Status(status)
	c.SetHeader("Content-Type", "text/html")
	c.Writer.Write([]byte(html))
}

func (c *Context) Data(status int, data []byte)  {
	c.Status(status)
	c.Writer.Write(data)
}
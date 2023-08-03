# github仓库

github地址：[ameamezhou/golang-web-frame](https://github.com/ameamezhou/golang-web-frame)
后续还将继续学习更新

## 创建github仓库
![在这里插入图片描述](https://img-blog.csdnimg.cn/5f2648e2321141f38137bd4fcb5ee13f.png)
设置免密登录
![在这里插入图片描述](https://img-blog.csdnimg.cn/e7aa9791ab7442039e7bde3a8fb00b2f.png)
ssh-keygen 一路回车就OK   上面有告诉你密钥生成地址
![在这里插入图片描述](https://img-blog.csdnimg.cn/3ba6a44946af4a11ad16984118f2d8dc.png)
红框为需要上传的公钥
![在这里插入图片描述](https://img-blog.csdnimg.cn/4ca0dc4e9270452e85a222a9302e1620.png)





# web框架学习


首先明确目标--**我们学习开发web框架的目的是** ：

在日常的web开发中，我们经常要使用到web框架，**python**就有很多好用的框架，比如**flask**和**django**，前者小巧精美，后者厚重却有着齐全的功能，不同开发者在设计框架的时候会有他们不同的看法和理念，因此在不同框架之间就会有许多不同的区别。这对于Go语言来说也是一样的，我们看到有很多好用的框架，例如**Beego**，**Gin**等等。但是我们在用这些框架的时候，我们可能需要去思考一下，其实这些框架翻找源码到底其实都是http等基础库构成的，但是我们为什么要使用它们呢？我们用框架究竟目的是什么？只有我们想明白了这一点我们才能更好的去做我们的开发工作，因此我决定做一个简单的框架实现这些基础功能。


# 开发周期
## 第一阶段--了解

```go
package main

import (
	"net/http"
)

func main()  {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe("localhost:9999", nil)
}


// 最基础的功能展示， 这里函数携带的参数是根据http库里面定义的
func sayHello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello world"))
}
```
实现一个最简单的web功能，就是打开页面输出 **hello world** 这里其实可以看到所需代码量其实和用现存的 gin 或者 beego 框架差不多，这里也能看出一些web框架大概的逻辑

然后我们在里面加点功能，增加点json输出

```go
package main

import (
	"encoding/json"
	"net/http"
)

func main()  {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe("localhost:9999", nil)
}


// 最基础的功能展示， 这里函数携带的参数是根据http库里面定义的
func sayHello(w http.ResponseWriter, r *http.Request){
	// 在页面输出展示json
	obj := make(map[string]interface{}, 0)
	obj["username"] = "xiawuyue"
	obj["password"] = "xiaoqizhou"

	// 这里是设置response 的响应头
	w.Header().Set("Content-Type", "application/json")
	// 这里是设置响应头的状态码  ok 就是 200
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(obj); err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write([]byte("hello world"))
}
```
然后我们可以看到页面
![在这里插入图片描述](https://img-blog.csdnimg.cn/4ae8bae7c96b40c78bfc308630259d89.png)
还是非常有意思的

## 第一阶段思考
两个问题
1 这个demo和你常用的框架的区别
2 你觉得这个地方的重点在哪里

附加：
关于web框架  我们都用过flask框架 请问这些框架最底层的运行逻辑是如何？go实现框架的逻辑相比于python如何？
（欢迎评论讨论）


### 小结
其实这一阶段我们要着重关注http的路由
```go
package main

import (
	"net/http"
)

func main()  {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe("localhost:9999", nil)
}


// 最基础的功能展示， 这里函数携带的参数是根据http库里面定义的
func sayHello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello world"))
}
```
在这里我们可以看到http.ListenAndServe 这里我们传进去的是一个nil，在里面是需要绑定路由的，也就是我们最关键的地方在HandleFunc这里，我们可以看到路由分发是通过 **http.HandleFunc("路径", 处理函数)** 这种形式实现的

```go
// HandleFunc registers the handler function for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

// HandleFunc registers the handler function for the given pattern.
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	if handler == nil {
		panic("http: nil handler")
	}
	mux.Handle(pattern, HandlerFunc(handler))
}
```
在ListenAndServe 这个函数中，我们第二个参数为nil，go会为我们分配一个默认的路由，会携带自己的**路由结构体** `ServeMux`
```go
// ListenAndServe always returns a non-nil error.
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}

// ServeMux还负责清除URL请求路径和主机标头，剥离端口号并重定向包含的任何请求。
// 或元素或重复的斜杠转换为等效的、更干净的URL。
type ServeMux struct {
	mu    sync.RWMutex // 这是一个互斥锁，保证并发
	m     map[string]muxEntry // 具体的路由规则
	es    []muxEntry // slice of entries sorted from longest to shortest.
	hosts bool       // whether any patterns contain hostnames 查看是否包含具体的host信息
}
```
其中我们最需要关注的就是这个m，我们注意到它是一个map类型，是一个 `string` 对应一个 `muxEntry` 结构体，这里最重要的就是`muxEntry`
Handler 其实就是一个interface接口，所以我们每一个HandFunc里面对应函数的类型都是要和这个 `ServeHTTP(ResponseWriter, *Request)` 保持一致的
```go
type muxEntry struct {
	h       Handler // 具体路由对应的 handler
	pattern string  // 匹配字符串
}

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```
题外话，go和python c 等语言不一样，这一块不需要通过sokcet来搞端口监听，http一个包就囊括了这些功能，所以我们可以深挖一下源码，看看究竟是怎么做得这方面的功能

## 第二阶段
看了上面的源码，我们实现的关键其实就是两个，一个是 `ServeMux` 一个是 `muxEntry` ，然后具体的 `Handler` 其实对应的就是一个**ServeHTTP**，我们需要实现的具体功能就是在这一块。所以其实我们完全可以自己来实现一个，不依赖 `net/http` 库它内置的一些功能，用我们自己的方式写一个 **ServeHTTP**

我们先梳理下这次的主要思路：

- base1的重点就是简单了解http库
- 我们来尝试自己写一个handle
- 以后我们的所有的框架代码都不再放在main.go 下面  养成包开发的习惯从主函数去调用

根据第一阶段的总结，我们不难发现我们要是想要自己实现一个框架，那么核心就是要实现一个 `muxEntry` 和 `Handler`
根据需求我们可以实现:
```go
package xiawuyue

import (
	"fmt"
	"net/http"
)

/*
本地包用法：
require xiawuyue v0.0.0

replace xiawuyue => ./base2/xiawuyue
*/

type XiaWuYue struct {
	router map[string]HandleFunc
}

// New 直接调用New方法构建对象
func New() *XiaWuYue {
	return &XiaWuYue{ router: make(map[string]HandleFunc) }
}

// HandleFunc 简单定义一类函数  这就是后续具体的处理方法的类型
type HandleFunc func(w http.ResponseWriter, req *http.Request)

func (x *XiaWuYue)addRoute(method string, pattern string, handleFunc HandleFunc) {
	// 其中method 是用来区分 get post 等方法的
	// patter 是提到的 muxEntry 中的匹配字符串 也就是具体的路径
	key := method + "-" + pattern
	x.router[key] = handleFunc
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

func (x *XiaWuYue) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Println("你访问的是根路径")
		w.Write([]byte("hello world"))
		// 这里会导致只在终端打印  所以要修改逻辑
	}

	key := req.Method + "-" + req.URL.Path
	if handler, ok := x.router[key]; ok {
		handler(w, req)
	}
	
}
```

这里Handler其实就是要求一个接口，这个接口它必须有 `ServerHTTP` 这个功能就ok，只要能理解这个，做这个逻辑的时候就会很清晰了，我们要实现的就是它的基本功能，并通过 `ServerHTTP` 对找到的路由提供相应的服务就行，所以这里我们新生成的 `struct xiawuyue` 它就需要带有这个功能接口

好的 代码看到这里我们来回忆一下第一天的内容：
```go
package main

import (
	"net/http"
)

func main()  {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe("localhost:9999", nil)
}


// 最基础的功能展示， 这里函数携带的参数是根据http库里面定义的
func sayHello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello world"))
}
```
这里我们看到原始的 `HandleFunc` 我们并没有初始化任何一个struct 对象，并且在 `ListenAndServe` 这里传进去的也是个 `nil` ， 这里的逻辑究竟是怎样的，我们为什么这样也能够去正常跑一个服务？

```go
// serverHandler delegates to either the server's Handler or
// DefaultServeMux and also handles "OPTIONS *" requests.
type serverHandler struct {
	srv *Server
}

func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
	handler := sh.srv.Handler
	if handler == nil {
		handler = DefaultServeMux
	}
	if req.RequestURI == "*" && req.Method == "OPTIONS" {
		handler = globalOptionsHandler{}
	}

	if req.URL != nil && strings.Contains(req.URL.RawQuery, ";") {
		var allowQuerySemicolonsInUse int32
		req = req.WithContext(context.WithValue(req.Context(), silenceSemWarnContextKey, func() {
			atomic.StoreInt32(&allowQuerySemicolonsInUse, 1)
		}))
		defer func() {
			if atomic.LoadInt32(&allowQuerySemicolonsInUse) == 0 {
				sh.srv.logf("http: URL query contains semicolon, which is no longer a supported separator; parts of the query may be stripped when parsed; see golang.org/issue/25192")
			}
		}()
	}

	handler.ServeHTTP(rw, req)
}
```
我们可以在 `http` 包的`server.go` 中找到这样一段，这里其实很好理解，当 `svr` 的 `handler` 为 `nil` 的时候，我们就会将 `DefaultServeMux` 导入当作这个 `muxEntry` ，它拥有 `ServerHTTP` 这个接口 可以实现相应的功能。

在我们都理解了这一块的知识之后，就写个总的 main.go 函数进行调用就ok

```go
package main

import (
	"fmt"
	"net/http"
	"xiawuyue/base2/xiawuyue"
)

// base1的重点就是简单了解http库
// 我们来尝试自己写一个handle

// 以后我们的所有的框架代码都不再放在main.go 下面  养成包开发的习惯
// 从主函数去调用

//func main(){
//	http.ListenAndServe(":9999", new(xiawuyue.XiaWuYue))
//}

func main()  {
	//如果没有 New 方法
	//r := new(xiawuyue.XiaWuYue)
	r := xiawuyue.New()
	r.Get("/get", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("hello world")
		w.Write([]byte("hello world get"))
		// 这里println 只会在终端输出  所以我们后续还是要包装一个w.return 的功能，其实很简单
	})
	// 请大家给斗鱼9999fg投一票 球球了
	http.ListenAndServe("localhost:9999", r)
	// 到这一步完成了然后就去启动
}
```
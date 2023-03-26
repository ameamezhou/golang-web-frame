# golang-web-frame
##### 这里所有的内容将会把过程和所有的经历体验同步更新到CSDN
##### 请关注 wx-zhou 嘻嘻

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
## 第一阶段思考
两个问题
1 这个demo和你常用的框架的区别
2 你觉得这个地方的重点在哪里

附加：
关于web框架  我们都用过flask框架 请问这些框架最底层的运行逻辑是如何？go实现框架的逻辑相比于python如何？


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



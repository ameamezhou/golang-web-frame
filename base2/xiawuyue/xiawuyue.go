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
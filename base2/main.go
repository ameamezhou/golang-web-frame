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
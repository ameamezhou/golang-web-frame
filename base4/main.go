package main

import (
	"fmt"
	xiawuyue "xiawuyue/base4/xiawuyue_base4"
)

func main()  {
	r := xiawuyue.New()
	r.Get("/", func(c *xiawuyue.Context) {
		c.HTML(200, "this is xiawuyue")
	})

	r.Get("/get", func(c *xiawuyue.Context) {
		c.HTML(200, "hello world")
	})
	// 有bug  修复完成
	r.Get("/hello/:name", func(c *xiawuyue.Context) {
		fmt.Println("test")
		c.String(200, "hello %s this is xiawuyue", c.Param("name"))
	})

	r.Get("/hello", func(c *xiawuyue.Context) {
		// test /hello?name=zhouzhougod
		c.String(200, "hello %s this is xiawuyue  query", c.Query("name"))
	})

	// 恭喜yyf领先一亿票拿下！
	r.Run("localhost:9999")

}

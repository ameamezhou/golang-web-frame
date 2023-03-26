package main

import (
	xiawuyue "xiawuyue/base3/xiawuyue_base3"
)

func main()  {
	r := xiawuyue.New()
	r.Get("/get", func(c *xiawuyue.Context) {
		c.HTML(200, "hello world")
	})
	// 恭喜yyf领先一亿票拿下！
	r.Run("localhost:9999")

}

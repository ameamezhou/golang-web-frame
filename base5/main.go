package main

import (
	"net/http"
	xiawuyue "xiawuyue/base5/xiawuyue_base5"
)

func main(){
	x := xiawuyue.New()
	x.Get("/test", func(c *xiawuyue.Context) {
		c.HTML(http.StatusOK, "<h1>test Page</h1>")
	})
	gOne := x.Group("/v1")
	{
		gOne.Get("/", func(c *xiawuyue.Context) {
			c.HTML(http.StatusOK, "<h1>V1  /</h1>")
		})

		gOne.Get("/hello", func(c *xiawuyue.Context) {
			c.HTML(http.StatusOK, "<h1>V1  hello</h1>")
		})
	}
	gTwo := x.Group("/v2")
	{
		gTwo.Get("/hello/:name", func(c *xiawuyue.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		gTwo.Post("/login", func(c *xiawuyue.Context) {
			c.Json(http.StatusOK, xiawuyue.Z{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	x.Run(":9999")
}
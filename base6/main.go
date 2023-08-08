package main

import (
	"net/http"
	xiawuyue "xiawuyue/base6/xiawuyue_base6"
)

func main() {
	r := xiawuyue.New()
	r.Use(xiawuyue.Logger()) // global midlleware
	r.Get("/", func(c *xiawuyue.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.Run(":9999")
}
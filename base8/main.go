package main

import (
	"net/http"
	xiawuyue "xiawuyue/base8/xiawuyue_base8"
)

func main() {
	r := xiawuyue.Default()
	r.Get("/", func(c *xiawuyue.Context) {
		c.String(http.StatusOK, "Hello xiawuyue\n")
	})
	r.Get("/panic", func(c *xiawuyue.Context) {
		names := []string{"xiawuyue"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
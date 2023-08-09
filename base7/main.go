package main

import xiawuyue "xiawuyue/base7/xiawuyue_base7"

func main()  {
	r := xiawuyue.New()
	r.Static("/assets", "/usr/xiawuyue/blog/static")

	r.Run(":9999")
}
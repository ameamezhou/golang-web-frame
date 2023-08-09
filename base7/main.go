package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	xiawuyue "xiawuyue/base7/xiawuyue_base7"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate() string {
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(t)
	return fmt.Sprintf("%s", t)
}

func main()  {
	r := xiawuyue.New()
	r.Use(xiawuyue.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "XiaWuYue", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.Get("/", func(c *xiawuyue.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.Get("/students", func(c *xiawuyue.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", xiawuyue.Z{
			"title":  "xiawuyue",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.Get("/date", func(c *xiawuyue.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", xiawuyue.Z{
			"title": "xiawuyue date",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
	//r := xiawuyue.New()
	//r.Static("/assets", "./static")
	//
	//r.Run(":9999")
}
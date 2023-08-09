### 模板渲染

#### 服务端渲染

现在越来越流行前后端分离的开发模式，即 web 后端提供 RestFul 接口，返回机构化的数据, 
通常为json或者xml,前端使用 Ajax 技术请求到所需数据, 利用 JavaScript 进行渲染。vue/react 等前端框架持续火热，
这种开发模式前后端解耦，优势非常突出，

前后端妃丽的一大问题就在于，页面是在客户端渲染的，比如浏览器，这对爬虫并不友好，Google爬虫已经能够爬取选然后的网页，但是短期内爬取服务器直接渲染HTML页面任是主流

今天介绍一下 web 框架如何支持服务器渲染的场景。

##### 静态文件 Serve Static Files

网页的三剑客，js，css，html。要做到服务器渲染，第一步就是要支持js、css等静态文件。
我们当时设计动态路由的时候，支持通配符*匹配多级子路径，比如路由规则 /assets/*filepath
可以匹配 /assets/开头的所有地址。例如 /assets/js/xiawuyue.js

匹配后filepath就赋值为 js/xiawuyue.js

那如果我们将所有的静态文件放在/usr/web目录下，那么filepath的值即是该目录下文件的相对地址.
映射到真实的文件后,将文件返回,静态服务器就实现了

##### HTML 模板渲染

Go语言内置了text/template和html/template2个模板标准库，其中html/template为 HTML 提供了较为完整的支持。包括普通变量渲染、列表渲染、对象渲染等



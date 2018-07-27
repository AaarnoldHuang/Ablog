package views

const blog  = `
# Martini  [![wercker status](https://app.wercker.com/status/174bef7e3c999e103cacfe2770102266 "wercker status")](https://app.wercker.com/project/bykey/174bef7e3c999e103cacfe2770102266) [![GoDoc](https://godoc.org/github.com/go-martini/martini?status.png)](http://godoc.org/github.com/go-martini/martini)

Martini是一个强大为了编写模块化Web应用而生的GO语言框架.

## 第一个应用

在你安装了GO语言和设置了你的[GOPATH](http://golang.org/doc/code.html#GOPATH)之后, 创建你的自己的.go文件, 这里我们假设它的名字叫做 server.go

~~~ go
package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Run()
}
~~~

然后安装Martini的包. (注意Martini需要Go语言1.1或者以上的版本支持):
~~~
go get github.com/go-martini/martini
~~~

最后运行你的服务:
~~~
go run server.go
~~~

这时你将会有一个Martini的服务监听了, 地址是: 

## 获得帮助

请加入: [邮件列表](https://groups.google.com/forum/#!forum/martini-go)

或者可以查看在线演示地址: [演示视频](http://martini.codegangsta.io/#demo)

## 功能列表
* 使用极其简单.
* 无侵入式的设计.
* 很好的与其他的Go语言包协同使用.
* 超赞的路径匹配和路由.
* 模块化的设计 - 容易插入功能件，也容易将其拔出来.
* 已有很多的中间件可以直接使用.
* 框架内已拥有很好的开箱即用的功能支持.
* **完全兼容[http.HandlerFunc](http://godoc.org/net/http#HandlerFunc)接口.**

## 更多中间件
更多的中间件和功能组件, 请查看代码仓库: [martini-contrib](https://github.com/martini-contrib).

## 目录
* [核心 Martini](#classic-martini)
  * [处理器](#handlers)
  * [路由](#routing)
  * [服务](#services)
  * [服务静态文件](#serving-static-files)
* [中间件处理器](#middleware-handlers)
  * [Next()](#next)
* [常见问答](#faq)

`

func GetBlog() string {
	return blog
}

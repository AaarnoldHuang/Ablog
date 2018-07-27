package main

import (
	"github.com/go-martini/martini"
	"gopkg.in/russross/blackfriday.v2"
	"./views"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string{
		return views.GetHead() + "<h2>空空如也~~~</h2>" + views.GetTail()
	})
	m.Get("/home", func()  string{
		return views.GetHead() + "<h2>空空如也~~~</h2>" + views.GetTail()
	})
	m.Get("/blogs", func() string{
		blog := string(blackfriday.Run([]byte(views.GetBlog())))
		return views.GetHead() + blog + views.GetTail()
	})
	m.Get("/about", func() string{
		return views.GetHead() + "<h2>我就是我</h2>" + views.GetTail()
	})
	m.Run()

}

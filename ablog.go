package main

import (
	"./controllers"
	"./models"
	"./views"
	"fmt"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martini-contrib/render"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var port string
	_, err := os.Stat("./config/config.toml")
	if err == nil {
		port = models.Config().Port
	} else {
		port = ":3000"
	}
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "htmls",
		Extensions: []string{".tmpl", ".html"},
	}))
	m.Get("/", func(r render.Render) string {
		if err != nil {
			dstFile, createrr := os.Create("./config/config.toml")
			if createrr != nil {
				fmt.Println(createrr)
			}
			defer dstFile.Close()
			return `<head><meta http-equiv="refresh" content="10"><meta http-equiv="refresh" content="0;url=/newconfig"></head>`

		}
		go models.ConnectDB()
		return `<head><meta http-equiv="refresh" content="10"><meta http-equiv="refresh" content="0;url=/home"></head>`
	})
	m.Get("/newconfig", func(r render.Render) {
		r.HTML(200, "newConfig", "")
	})

	m.Post("/newconfig", func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		Title := req.Form["Title"][0]
		Port := req.Form["Port"][0]
		Username := req.Form["UserName"][0]
		Password := req.Form["Password"][0]
		Bio := req.Form["Bio"][0]
		DBname := req.Form["DBname"][0]
		DBuser := req.Form["DBuser"][0]
		DBpwd := req.Form["DBpasswd"][0]
		config := []byte(controllers.Writeconfig(Title, Port, Username, Password, Bio, DBname, DBuser, DBpwd))
		err := ioutil.WriteFile("./config/config.toml", config, 0644)
		if err != nil {
			fmt.Println(err)
		}
	})
	m.Get("/home", func(r render.Render) {
		r.HTML(200, "blogs", template.HTML("<h2>空空如也~~~</h2>"))
	})
	m.Get("/blogs", func(r render.Render) {
		blog := string(blackfriday.Run([]byte(views.GetBlog())))
		r.HTML(200, "postView", controllers.PostView("Test Blogs", template.HTML(blog)))
	})
	m.Get("/about", func(r render.Render) {
		models.WriteBlogtoDB("'测试-Martini介绍'", views.GetBlog(), "'published'", "''")
		r.HTML(200, "blogs", "我就是我")
	})
	m.RunOnAddr(port)
	m.Run()

}

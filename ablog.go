package main

import (
	"./controllers"
	"./models"
	"fmt"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martini-contrib/render"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"io/ioutil"
	"net/http"
	"sort"
)

type getPost struct {
	Id    string
	Title string
	Body  template.HTML
}
type postView struct {
	Title string
	Body  template.HTML
}

func main() {
	var port string
	if models.Config().First {
		port = ":3000"
	} else {
		port = models.Config().Port
	}
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "htmls",
		Extensions: []string{".tmpl", ".html"},
	}))
	m.Get("/", func(r render.Render) string {
		if models.Config().First {
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
		r.HTML(200, "blogs", template.HTML(`<h4>BIENVENIDO A GALLIFREY!<h4></br>
		<h4>欢迎来到我的Gallifrey！</h4></br>`))
	})
	m.Get("/blogs", func(r render.Render) {
		blog := ""
		var keys []int
		posts, result := models.GetAllBlogsfromDB()
		if !result {
			r.HTML(200, "blogs", "An error occurs while trying to connect to the database. Please contact your administrator！")
		}
		for k := range posts {
			keys = append(keys, k)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
		for _, k := range keys {
			blog += fmt.Sprintf("<li> <a class=\"active\" href=\"/blogs/%s\">%s</a> </li>  \n", posts[k].Id, posts[k].Title)
		}

		r.HTML(200, "blogs", template.HTML(blog))
	})

	m.Get("/blogs/**", func(params martini.Params, r render.Render) {
		id := params["_1"]
		if id == "blogs" {
			r.HTML(200, "jump", "/blogs")
		} else if id == "home" {
			r.HTML(200, "jump", "/home")
		} else if id == "about" {
			r.HTML(200, "jump", "/about")
		}
		blog := models.GetABlogfromDB(id)
		post := postView{
			Title: blog.Title,
			Body:  template.HTML(blackfriday.Run([]byte(blog.Body))),
		}
		r.HTML(200, "postView", post)
	})
	m.Get("/about", func(r render.Render) {
		r.HTML(200, "blogs", "我就是我")
	})
	m.RunOnAddr(port)
	m.Run()

}

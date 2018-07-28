package main

import (
	"./controllers"
	"./views"
	"fmt"
	"github.com/go-martini/martini"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var port string
	_, err := os.Stat("./config/config.toml")
	if err == nil {
		port = controllers.Config().Port
	} else {
		port = ":3000"
	}
	m := martini.Classic()
	m.Get("/", func() string {
		if err != nil {
			dstFile, createrr := os.Create("./config/config.toml")
			if createrr != nil {
				fmt.Println(createrr)
			}
			defer dstFile.Close()
			return `<head><meta http-equiv="refresh" content="10"><meta http-equiv="refresh" content="0;url=/newconfig"></head>`

		}
		return views.GetHead() + "<h2>空空如也~~~</h2>" + views.GetTail()
	})
	m.Get("/newconfig", func() string {
		return views.GetNewConfig()
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
	m.Get("/home", func() string {
		return views.GetHead() + "<h2>空空如也~~~</h2>" + views.GetTail()
	})
	m.Get("/blogs", func() string {
		blog := string(blackfriday.Run([]byte(views.GetBlog())))
		return views.GetHead() + blog + views.GetTail()
	})
	m.Get("/about", func() string {
		return views.GetHead() + "<h2>我就是我</h2>" + views.GetTail()
	})
	m.RunOnAddr(port)
	m.Run()

}

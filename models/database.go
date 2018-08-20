package models

import (
	"database/sql"
	"fmt"
	"html/template"
)

type getPost struct {
	Id    string
	Title string
	Body  template.HTML
}

type getBlog struct {
	Title string
	Body  string
}

const creattable string = "CREATE TABLE ablog_post (`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT," +
	"`post_date` DATETIME NOT NULL DEFAULT NOW(),`post_content` LONGTEXT NOT NULL,`post_title`  TEXT NOT NULL," +
	"`post_name` VARCHAR(200) NOT NULL default '',`post_status` VARCHAR(20) NOT NULL default 'publish'," +
	"`post_password` VARCHAR(255) NOT NULL default '',`post_modified` DATETIME NOT NULL default NOW()," +
	"PRIMARY KEY (`id`));"

var dbinfo = fmt.Sprintf("%s:%s@/%s?charset=utf8", Config().DB.DBuser, Config().DB.DBpasswd, Config().DB.DBname)

func ConnectDB() {
	var db *sql.DB
	db, _ = sql.Open("mysql", dbinfo)
	table, err := db.Query("SELECT * FROM ablog_post")
	if table == nil && err != nil {
		_, errc := db.Query(creattable)
		if errc != nil {
			fmt.Println(errc)
		}
	}
	defer db.Close()
}
func WriteBlogtoDB(title, content, status, passwd string) bool {
	db, _ := sql.Open("mysql", dbinfo)
	cmd := fmt.Sprintf("INSERT ablog_post SET post_title=%s,post_content=%s,post_status=%s,post_password=%s;", title, content, status, passwd)
	res, err := db.Exec(cmd)
	if err != nil {
		fmt.Println(err)
	}
	id, _ := res.LastInsertId()
	fmt.Printf("[Database] Insert ID: %d successd \n", id)
	defer db.Close()
	return true
}

func GetABlogfromDB(id string) getBlog {
	var post getBlog
	db, _ := sql.Open("mysql", dbinfo)
	res, err := db.Query(fmt.Sprintf("select * from ablog_post where id='%s';", id))
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() {
		var id string
		var post_title string
		var post_content string
		var post_date string
		var post_status string
		var post_name string
		var post_password string
		var post_modified string
		errr := res.Scan(&id, &post_date, &post_content, &post_title, &post_name, &post_status, &post_password, &post_modified)
		if errr != nil {
			fmt.Println(errr)
		}
		post.Title = post_title
		post.Body = post_content
	}
	defer db.Close()
	return post
}

func GetAllBlogsfromDB() (map[int]getPost, bool) {
	i := 0
	posts := make(map[int]getPost)
	db, dberr := sql.Open("mysql", dbinfo)
	if dberr != nil {
		return posts, false
	}
	res, err := db.Query("select * from ablog_post where post_status='published';")
	if err != nil {
		return posts, false
	}
	for res.Next() {
		var post getPost
		var id string
		var post_title string
		var post_content string
		var post_date string
		var post_status string
		var post_name string
		var post_password string
		var post_modified string
		errr := res.Scan(&id, &post_date, &post_content, &post_title, &post_name, &post_status, &post_password, &post_modified)
		if errr != nil {
			return posts, false
		}
		post.Id = id
		post.Title = post_title
		post.Body = template.HTML(post_content)
		posts[i] = post
		i++
	}
	defer db.Close()
	return posts, true
}

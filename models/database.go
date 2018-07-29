package models

import (
	"database/sql"
	"fmt"
)

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

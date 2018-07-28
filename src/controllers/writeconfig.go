package controllers

import "fmt"

func Writeconfig(title, port, username, password, bio, dbname, dbuser, dbpasswd string) string {
	return fmt.Sprintf("title = \"%s\" \nport = \":%s\" \n[admin]\nname = \"%s\"\npasswd = \"%s\"\nbio = \"%s\"\n"+
		"[database]\nserver = \"127.0.0.1\"\nports = [ 3306 ]\ndbname = \"%s\"\ndbuser = \"%s\"\ndbpasswd = \"%s\"\n"+
		"connection_max = 5000\nenabled = true", title, port, username, password, bio, dbname, dbuser, dbpasswd)
}

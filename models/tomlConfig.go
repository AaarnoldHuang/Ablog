package models

type tomlconfig struct {
	Title string
	Port  string
	Admin admin    `toml:"admin"`
	DB    database `toml:"database"`
}

type admin struct {
	Name   string
	Passwd string
	Bio    string
}

type database struct {
	Server   string
	Ports    []int
	DBname   string
	DBuser   string
	DBpasswd string
	ConnMax  int `toml:"connection_max"`
	Enabled  bool
}

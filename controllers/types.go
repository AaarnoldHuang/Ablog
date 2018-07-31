package controllers

import "html/template"

type postView struct {
	Title string
	Body  template.HTML
}

type getPost struct {
	Id    string
	Title string
	Body  template.HTML
}

func PostView(title string, body template.HTML) postView {
	return postView{
		Title: title,
		Body:  body,
	}
}

func GetPost(id, title string, body template.HTML) getPost {
	return getPost{
		Id:    id,
		Title: title,
		Body:  body,
	}
}

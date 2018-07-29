package controllers

import "html/template"

type postView struct {
	Title string
	Body  template.HTML
}

func PostView(title string, body template.HTML) postView {
	return postView{
		Title: title,
		Body:  body,
	}
}

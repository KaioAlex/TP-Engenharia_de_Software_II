package controllers

import "ll/views"

func NewStatic() *Static {
	return &Static{
		Home: views.NewView("bootstrap", "static/home"),
	}
}

type Static struct {
	Home *views.View
}

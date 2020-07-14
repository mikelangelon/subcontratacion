package web

import (
	"html/template"
	"net/http"
)

type ProfileResource struct {
}

type profile struct {
	Profile *Profile
}

func (r ProfileResource) Profile(resp http.ResponseWriter, req *http.Request) {
	p, _ := LoginResource{}.Auth(resp, req)
	t, err := template.ParseFiles(
		"web/template/layout.html",
		"web/template/topmenu.html",
		"web/template/leftmenu.html",
		"web/template/budgetrequests.html",
		"web/template/banner.html",
		"web/template/providers.html",
		"web/content/profile.html")
	if err != nil {
		//TODO
	}
	prof := profile{
		Profile: p,
	}
	t.Execute(resp, prof)
}

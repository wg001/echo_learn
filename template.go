package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type Template struct {
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if t, ok := templates[name]; ok {
		return t.ExecuteTemplate(w, "layout.html", data)
	}
	c.Echo().Logger.Debugf("Template[%s] Not Found.", name)
	return templates["error"].ExecuteTemplate(w, "layout.html", "Internal Server Error")
}

func LoadTemplates() {
	var baseTemplate = "public/views/layout.html"
	templates = make(map[string]*template.Template)
	// 各HTMLテンプレートに共通レイアウトを適用した結果をmapに保存する
	templates["index"] = template.Must(
		template.ParseFiles(baseTemplate, "public/views/index.html"))
	templates["error"] = template.Must(
		template.ParseFiles(baseTemplate, "public/views/error.html"))
	templates["user"] = template.Must(
		template.ParseFiles(baseTemplate, "public/views/user.html"))
	templates["login"] = template.Must(
		template.ParseFiles(baseTemplate, "public/views/login.html"))
	templates["admin"] = template.Must(
		template.ParseFiles(baseTemplate, "public/views/admin.html"))
	templates["admin_users"] = template.Must(
		template.ParseFiles(baseTemplate, "public/views/admin_users.html"))
}

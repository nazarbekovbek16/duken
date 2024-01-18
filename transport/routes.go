package transport

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (s *Server) routes() {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	s.HTTP.Static("/", "templates")

	s.HTTP.Static("/css", "templates/css")

	s.HTTP.Renderer = renderer

	s.HTTP.GET("/", s.h.Home)

	s.HTTP.GET("/login", s.h.User.LoginPage)
	s.HTTP.POST("/login", s.h.User.Login)

	s.HTTP.GET("/register", s.h.User.Registration)
	s.HTTP.POST("/register", s.h.User.Register)
	s.HTTP.GET("/logout", s.h.User.Logout)
}
